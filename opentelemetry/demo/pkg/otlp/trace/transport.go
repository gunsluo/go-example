package trace

import (
	"net/http"

	"go.opentelemetry.io/otel/api/kv"
	"go.opentelemetry.io/otel/api/standard"
	"go.opentelemetry.io/otel/api/trace"
	"go.opentelemetry.io/otel/plugin/httptrace"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
)

// Transport wraps a RoundTripper. If a request is being traced with
// Tracer, Transport will inject the current span into the headers,
// and set HTTP related tags on the span.
type Transport struct {
	// The actual RoundTripper to use for the request. A nil
	// RoundTripper defaults to http.DefaultTransport.
	http.RoundTripper

	tracer        trace.Tracer
	enable        bool
	componentName string
	logger        *zap.Logger
	opNameFunc    func(*http.Request) string

	disableInjectSpanContext bool
}

type TransportOption func(*Transport)

func WithTransportComponentName(componentName string) TransportOption {
	return func(t *Transport) {
		t.componentName = componentName
	}
}

func WithTransportOpNameFunc(opNameFunc func(*http.Request) string) TransportOption {
	return func(t *Transport) {
		t.opNameFunc = opNameFunc
	}
}

func WithTransportLogger(logger *zap.Logger) TransportOption {
	return func(t *Transport) {
		t.logger = logger
	}
}

func NewTransport(tracer trace.Tracer, options ...TransportOption) (*Transport, error) {
	t := &Transport{
		componentName: "net/http",
		logger:        zap.NewNop(),
		opNameFunc: func(r *http.Request) string {
			return r.Method + " " + r.URL.String() + " - HTTP Client"
		},
	}

	for _, opt := range options {
		opt(t)
	}

	if tracer == nil {
		return t, nil
	}

	if _, ok := tracer.(*trace.NoopTracer); ok {
		return t, nil
	}

	t.tracer = tracer
	t.enable = true

	t.logger.With(zap.String("component", t.componentName)).Info("success to enable trace http transport")
	return t, nil
}

// RoundTrip implements the RoundTripper interface.
func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	rt := t.RoundTripper
	if rt == nil {
		rt = http.DefaultTransport
	}

	if !t.enable {
		return rt.RoundTrip(req)
	}

	ctx, span := t.tracer.Start(req.Context(), t.opNameFunc(req))
	span.SetAttributes(
		kv.Key("component").String(t.componentName),
		standard.HTTPUrlKey.String(req.URL.String()),
		standard.HTTPSchemeKey.String(req.URL.Scheme),
		standard.HTTPMethodKey.String(req.Method),
	)
	defer span.End()

	httptrace.Inject(ctx, req)

	resp, err := rt.RoundTrip(req)
	if err != nil {
		span.SetStatus(codes.Internal, err.Error())
		return resp, err
	}
	if resp.StatusCode >= http.StatusBadRequest {
		span.SetStatus(codes.Internal, "invalid code")
	} else {
		span.SetAttributes(
			standard.HTTPStatusCodeKey.Int(resp.StatusCode),
		)
	}

	return resp, nil
}
