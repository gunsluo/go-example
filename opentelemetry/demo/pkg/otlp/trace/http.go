package trace

import (
	"io"
	"net/http"

	"go.opentelemetry.io/otel/api/correlation"
	"go.opentelemetry.io/otel/api/kv"
	"go.opentelemetry.io/otel/api/standard"
	"go.opentelemetry.io/otel/api/trace"
	"google.golang.org/grpc/codes"

	//"go.opentelemetry.io/otel/instrumentation/httptrace"
	"go.opentelemetry.io/otel/plugin/httptrace"
	"go.uber.org/zap"
)

// HttpMiddleware is http middleware about trace
type HttpMiddleware struct {
	tracer        trace.Tracer
	enable        bool
	componentName string
	opNameFunc    func(r *http.Request) string
	logger        *zap.Logger
}

const defaultHttpComponentName = "net/http"

type HttpOption func(*HttpMiddleware)

func NewHttpMiddleware(tracer trace.Tracer, opts ...HttpOption) (*HttpMiddleware, error) {
	m := &HttpMiddleware{
		tracer:        tracer,
		enable:        true,
		opNameFunc:    defaultOpNameFunc,
		componentName: defaultHttpComponentName}
	for _, opt := range opts {
		opt(m)
	}

	return m, nil
}

// WithHttpEnable disable to collect trace
func WithHttpEnable(enable bool) HttpOption {
	return func(m *HttpMiddleware) {
		m.enable = enable
	}
}

// WithHttpComponentName set component name
func WithHttpComponentName(componentName string) HttpOption {
	return func(m *HttpMiddleware) {
		m.componentName = componentName
	}
}

// WithHttpOpNameFunc set function to get operation name for span
func WithHttpOpNameFunc(fn func(r *http.Request) string) HttpOption {
	return func(m *HttpMiddleware) {
		m.opNameFunc = fn
	}
}

// WithHttpLogger set logger
func WithHttpLogger(logger *zap.Logger) HttpOption {
	return func(m *HttpMiddleware) {
		m.logger = logger
	}
}

// defaultOpNameFunc is default function that get operation name from http request
func defaultOpNameFunc(r *http.Request) string {
	return r.URL.Scheme + " " + r.Method + " " + r.URL.Path
}

func (m *HttpMiddleware) Handle(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !m.enable {
			h(w, r)
			return
		}

		attrs, entries, spanCtx := httptrace.Extract(r.Context(), r)
		ctx, span := m.tracer.Start(
			trace.ContextWithRemoteSpanContext(r.Context(), spanCtx),
			m.opNameFunc(r),
			trace.WithAttributes(attrs...),
		)

		r = r.WithContext(correlation.ContextWithMap(ctx, correlation.NewMap(correlation.MapUpdate{
			MultiKV: entries,
		})))

		sct := &statusCodeWriterTracker{ResponseWriter: w}
		defer func() {
			span.SetAttributes(
				kv.Key("component").String(m.componentName),
				standard.HTTPStatusCodeKey.Int(sct.status),
				standard.HTTPSchemeKey.String(r.URL.Scheme),
				standard.HTTPMethodKey.String(r.Method),
				//standard.HTTPHostKey.String(r.Host),
				//standard.HTTPClientIPKey.String(r.Host),
			)

			if sct.status >= http.StatusInternalServerError || !sct.wroteheader {
				span.SetStatus(codes.Internal, "invalid code")
				//span.SetAttributes(kv.Key("error").Bool(true))
			}
			span.End()
		}()

		h.ServeHTTP(sct.wrappedResponseWriter(), r)
	})
}

type statusCodeWriterTracker struct {
	http.ResponseWriter
	status      int
	wroteheader bool
}

func (w *statusCodeWriterTracker) WriteHeader(status int) {
	w.status = status
	w.wroteheader = true
	w.ResponseWriter.WriteHeader(status)
}

func (w *statusCodeWriterTracker) Write(b []byte) (int, error) {
	if !w.wroteheader {
		w.wroteheader = true
		w.status = 200
	}
	return w.ResponseWriter.Write(b)
}

// wrappedResponseWriter returns a wrapped version of the original
// ResponseWriter and only implements the same combination of additional
// interfaces as the original.  This implementation is based on
// https://github.com/felixge/httpsnoop.
func (w *statusCodeWriterTracker) wrappedResponseWriter() http.ResponseWriter {
	var (
		hj, i0 = w.ResponseWriter.(http.Hijacker)
		cn, i1 = w.ResponseWriter.(http.CloseNotifier)
		pu, i2 = w.ResponseWriter.(http.Pusher)
		fl, i3 = w.ResponseWriter.(http.Flusher)
		rf, i4 = w.ResponseWriter.(io.ReaderFrom)
	)

	switch {
	case !i0 && !i1 && !i2 && !i3 && !i4:
		return struct {
			http.ResponseWriter
		}{w}
	case !i0 && !i1 && !i2 && !i3 && i4:
		return struct {
			http.ResponseWriter
			io.ReaderFrom
		}{w, rf}
	case !i0 && !i1 && !i2 && i3 && !i4:
		return struct {
			http.ResponseWriter
			http.Flusher
		}{w, fl}
	case !i0 && !i1 && !i2 && i3 && i4:
		return struct {
			http.ResponseWriter
			http.Flusher
			io.ReaderFrom
		}{w, fl, rf}
	case !i0 && !i1 && i2 && !i3 && !i4:
		return struct {
			http.ResponseWriter
			http.Pusher
		}{w, pu}
	case !i0 && !i1 && i2 && !i3 && i4:
		return struct {
			http.ResponseWriter
			http.Pusher
			io.ReaderFrom
		}{w, pu, rf}
	case !i0 && !i1 && i2 && i3 && !i4:
		return struct {
			http.ResponseWriter
			http.Pusher
			http.Flusher
		}{w, pu, fl}
	case !i0 && !i1 && i2 && i3 && i4:
		return struct {
			http.ResponseWriter
			http.Pusher
			http.Flusher
			io.ReaderFrom
		}{w, pu, fl, rf}
	case !i0 && i1 && !i2 && !i3 && !i4:
		return struct {
			http.ResponseWriter
			http.CloseNotifier
		}{w, cn}
	case !i0 && i1 && !i2 && !i3 && i4:
		return struct {
			http.ResponseWriter
			http.CloseNotifier
			io.ReaderFrom
		}{w, cn, rf}
	case !i0 && i1 && !i2 && i3 && !i4:
		return struct {
			http.ResponseWriter
			http.CloseNotifier
			http.Flusher
		}{w, cn, fl}
	case !i0 && i1 && !i2 && i3 && i4:
		return struct {
			http.ResponseWriter
			http.CloseNotifier
			http.Flusher
			io.ReaderFrom
		}{w, cn, fl, rf}
	case !i0 && i1 && i2 && !i3 && !i4:
		return struct {
			http.ResponseWriter
			http.CloseNotifier
			http.Pusher
		}{w, cn, pu}
	case !i0 && i1 && i2 && !i3 && i4:
		return struct {
			http.ResponseWriter
			http.CloseNotifier
			http.Pusher
			io.ReaderFrom
		}{w, cn, pu, rf}
	case !i0 && i1 && i2 && i3 && !i4:
		return struct {
			http.ResponseWriter
			http.CloseNotifier
			http.Pusher
			http.Flusher
		}{w, cn, pu, fl}
	case !i0 && i1 && i2 && i3 && i4:
		return struct {
			http.ResponseWriter
			http.CloseNotifier
			http.Pusher
			http.Flusher
			io.ReaderFrom
		}{w, cn, pu, fl, rf}
	case i0 && !i1 && !i2 && !i3 && !i4:
		return struct {
			http.ResponseWriter
			http.Hijacker
		}{w, hj}
	case i0 && !i1 && !i2 && !i3 && i4:
		return struct {
			http.ResponseWriter
			http.Hijacker
			io.ReaderFrom
		}{w, hj, rf}
	case i0 && !i1 && !i2 && i3 && !i4:
		return struct {
			http.ResponseWriter
			http.Hijacker
			http.Flusher
		}{w, hj, fl}
	case i0 && !i1 && !i2 && i3 && i4:
		return struct {
			http.ResponseWriter
			http.Hijacker
			http.Flusher
			io.ReaderFrom
		}{w, hj, fl, rf}
	case i0 && !i1 && i2 && !i3 && !i4:
		return struct {
			http.ResponseWriter
			http.Hijacker
			http.Pusher
		}{w, hj, pu}
	case i0 && !i1 && i2 && !i3 && i4:
		return struct {
			http.ResponseWriter
			http.Hijacker
			http.Pusher
			io.ReaderFrom
		}{w, hj, pu, rf}
	case i0 && !i1 && i2 && i3 && !i4:
		return struct {
			http.ResponseWriter
			http.Hijacker
			http.Pusher
			http.Flusher
		}{w, hj, pu, fl}
	case i0 && !i1 && i2 && i3 && i4:
		return struct {
			http.ResponseWriter
			http.Hijacker
			http.Pusher
			http.Flusher
			io.ReaderFrom
		}{w, hj, pu, fl, rf}
	case i0 && i1 && !i2 && !i3 && !i4:
		return struct {
			http.ResponseWriter
			http.Hijacker
			http.CloseNotifier
		}{w, hj, cn}
	case i0 && i1 && !i2 && !i3 && i4:
		return struct {
			http.ResponseWriter
			http.Hijacker
			http.CloseNotifier
			io.ReaderFrom
		}{w, hj, cn, rf}
	case i0 && i1 && !i2 && i3 && !i4:
		return struct {
			http.ResponseWriter
			http.Hijacker
			http.CloseNotifier
			http.Flusher
		}{w, hj, cn, fl}
	case i0 && i1 && !i2 && i3 && i4:
		return struct {
			http.ResponseWriter
			http.Hijacker
			http.CloseNotifier
			http.Flusher
			io.ReaderFrom
		}{w, hj, cn, fl, rf}
	case i0 && i1 && i2 && !i3 && !i4:
		return struct {
			http.ResponseWriter
			http.Hijacker
			http.CloseNotifier
			http.Pusher
		}{w, hj, cn, pu}
	case i0 && i1 && i2 && !i3 && i4:
		return struct {
			http.ResponseWriter
			http.Hijacker
			http.CloseNotifier
			http.Pusher
			io.ReaderFrom
		}{w, hj, cn, pu, rf}
	case i0 && i1 && i2 && i3 && !i4:
		return struct {
			http.ResponseWriter
			http.Hijacker
			http.CloseNotifier
			http.Pusher
			http.Flusher
		}{w, hj, cn, pu, fl}
	case i0 && i1 && i2 && i3 && i4:
		return struct {
			http.ResponseWriter
			http.Hijacker
			http.CloseNotifier
			http.Pusher
			http.Flusher
			io.ReaderFrom
		}{w, hj, cn, pu, fl, rf}
	default:
		return struct {
			http.ResponseWriter
		}{w}
	}
}
