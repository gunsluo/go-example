package trace

import (
	"context"
	"net"
	"regexp"

	"go.opentelemetry.io/otel/api/correlation"
	"go.opentelemetry.io/otel/api/global"
	"go.opentelemetry.io/otel/api/kv"
	"go.opentelemetry.io/otel/api/propagation"
	"go.opentelemetry.io/otel/api/standard"
	"go.opentelemetry.io/otel/api/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

// UnaryServerInterceptor returns a grpc.UnaryServerInterceptor suitable
// for use in a grpc.NewServer call.
//
// For example:
//     tracer := global.Tracer("client-tracer")
//     s := grpc.Dial(
//         grpc.UnaryInterceptor(grpctrace.UnaryServerInterceptor(tracer)),
//         ...,  // (existing ServerOptions))
func UnaryServerInterceptor(tracer trace.Tracer, componentName string) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		requestMetadata, _ := metadata.FromIncomingContext(ctx)
		metadataCopy := requestMetadata.Copy()

		entries, spanCtx := Extract(ctx, &metadataCopy)
		ctx = correlation.ContextWithMap(ctx, correlation.NewMap(correlation.MapUpdate{
			MultiKV: entries,
		}))

		ctx, span := tracer.Start(
			trace.ContextWithRemoteSpanContext(ctx, spanCtx),
			info.FullMethod,
			trace.WithSpanKind(trace.SpanKindServer),
			trace.WithAttributes(kv.Key("component").String(componentName)),
			trace.WithAttributes(peerInfoFromContext(ctx)...),
			trace.WithAttributes(standard.RPCServiceKey.String(serviceFromFullMethod(info.FullMethod))),
		)
		defer span.End()

		//messageReceived.Event(ctx, 1, req)
		resp, err := handler(ctx, req)
		if err != nil {
			s, _ := status.FromError(err)
			span.SetStatus(s.Code(), s.Message())
			//messageSent.Event(ctx, 1, s.Proto())
		} else {
			//messageSent.Event(ctx, 1, resp)
		}

		return resp, err
	}
}

// UnaryClientInterceptor returns a grpc.UnaryClientInterceptor suitable
// for use in a grpc.Dial call.
//
// For example:
//     tracer := global.Tracer("client-tracer")
//     s := grpc.NewServer(
//         grpc.WithUnaryInterceptor(grpctrace.UnaryClientInterceptor(tracer)),
//         ...,  // (existing DialOptions))
func UnaryClientInterceptor(tracer trace.Tracer, componentName string) grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		requestMetadata, _ := metadata.FromOutgoingContext(ctx)
		metadataCopy := requestMetadata.Copy()

		var span trace.Span
		ctx, span = tracer.Start(
			ctx, method+" - GRPC Client",
			trace.WithSpanKind(trace.SpanKindClient),
			trace.WithAttributes(kv.Key("component").String(componentName)),
			trace.WithAttributes(peerInfoFromTarget(cc.Target())...),
			trace.WithAttributes(standard.RPCServiceKey.String(serviceFromFullMethod(method))),
		)
		defer span.End()

		Inject(ctx, &metadataCopy)
		ctx = metadata.NewOutgoingContext(ctx, metadataCopy)

		//messageSent.Event(ctx, 1, req)
		err := invoker(ctx, method, req, reply, cc, opts...)
		//messageReceived.Event(ctx, 1, reply)
		if err != nil {
			s, _ := status.FromError(err)
			span.SetStatus(s.Code(), s.Message())
		}

		return err
	}
}

// Inject injects correlation context and span context into the gRPC
// metadata object. This function is meant to be used on outgoing
// requests.
func Inject(ctx context.Context, metadata *metadata.MD, opts ...GRPCOption) {
	c := newConfig(opts)
	propagation.InjectHTTP(ctx, c.propagators, &metadataSupplier{
		metadata: metadata,
	})
}

// Extract returns the correlation context and span context that
// another service encoded in the gRPC metadata object with Inject.
// This function is meant to be used on incoming requests.
func Extract(ctx context.Context, metadata *metadata.MD, opts ...GRPCOption) ([]kv.KeyValue, trace.SpanContext) {
	c := newConfig(opts)
	ctx = propagation.ExtractHTTP(ctx, c.propagators, &metadataSupplier{
		metadata: metadata,
	})

	spanContext := trace.RemoteSpanContextFromContext(ctx)
	var correlationCtxKVs []kv.KeyValue
	correlation.MapFromContext(ctx).Foreach(func(kv kv.KeyValue) bool {
		correlationCtxKVs = append(correlationCtxKVs, kv)
		return true
	})

	return correlationCtxKVs, spanContext
}

func peerInfoFromTarget(target string) []kv.KeyValue {
	host, port, err := net.SplitHostPort(target)

	if err != nil {
		return []kv.KeyValue{}
	}

	if host == "" {
		host = "127.0.0.1"
	}

	return []kv.KeyValue{
		standard.NetPeerIPKey.String(host),
		standard.NetPeerPortKey.String(port),
	}
}

func peerInfoFromContext(ctx context.Context) []kv.KeyValue {
	p, ok := peer.FromContext(ctx)

	if !ok {
		return []kv.KeyValue{}
	}

	return peerInfoFromTarget(p.Addr.String())
}

var fullMethodRegexp = regexp.MustCompile(`^\/?(?:\S+\.)?(\S+)\/\S+$`)

func serviceFromFullMethod(method string) string {
	match := fullMethodRegexp.FindStringSubmatch(method)
	if len(match) == 0 {
		return ""
	}

	return match[1]
}

// GRPCOption is a function that allows configuration of the grpctrace Extract()
// and Inject() functions
type GRPCOption func(*config)

type config struct {
	propagators propagation.Propagators
}

func newConfig(opts []GRPCOption) *config {
	c := &config{propagators: global.Propagators()}
	for _, o := range opts {
		o(c)
	}
	return c
}

// WithPropagators sets the propagators to use for Extraction and Injection
func WithPropagators(props propagation.Propagators) GRPCOption {
	return func(c *config) {
		c.propagators = props
	}
}

type metadataSupplier struct {
	metadata *metadata.MD
}

func (s *metadataSupplier) Get(key string) string {
	values := s.metadata.Get(key)
	if len(values) == 0 {
		return ""
	}
	return values[0]
}

func (s *metadataSupplier) Set(key string, value string) {
	s.metadata.Set(key, value)
}
