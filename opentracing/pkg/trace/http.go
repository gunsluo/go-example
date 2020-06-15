package trace

import (
	"fmt"
	"net/http"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

type HttpMiddleware struct {
	tracer opentracing.Tracer
	opts   httpOptions
}

type HttpOption interface {
	apply(*httpOptions)
}

type fnHttpOption struct {
	fn func(*httpOptions)
}

func (o *fnHttpOption) apply(opts *httpOptions) {
	o.fn(opts)
}

// WithHttpEnable disable to collect trace
func WithHttpEnable(enable bool) HttpOption {
	return &fnHttpOption{fn: func(opts *httpOptions) {
		opts.enable = enable
	}}
}

// WithHttpComponentName set component name
func WithHttpComponentName(componentName string) HttpOption {
	return &fnHttpOption{fn: func(opts *httpOptions) {
		opts.componentName = componentName
	}}
}

// WithOpNameFunc set function to get operation name for span
func WithOpNameFunc(fn func(r *http.Request) string) HttpOption {
	return &fnHttpOption{fn: func(opts *httpOptions) {
		opts.opNameFunc = fn
	}}
}

type httpOptions struct {
	enable        bool
	componentName string
	opNameFunc    func(r *http.Request) string
}

const defaultComponentName = "net/http"

func defaultHttpOptions() httpOptions {
	return httpOptions{
		enable:     true,
		opNameFunc: defaultOpNameFunc}
}

func defaultOpNameFunc(r *http.Request) string {
	return r.URL.Scheme + " " + r.Method + " " + r.URL.Path
}

func NewHttpMiddleware(tracer opentracing.Tracer, opts ...HttpOption) *HttpMiddleware {
	m := &HttpMiddleware{tracer: tracer, opts: defaultHttpOptions()}
	for _, opt := range opts {
		opt.apply(&m.opts)
	}

	return m
}

func (m *HttpMiddleware) Handle(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !m.opts.enable {
			h(w, r)
			return
		}

		ctx, err := m.tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
		fmt.Println("--->", ctx, err)

		//if span := opentracing.SpanFromContext(ctx); span != nil {

		sp := m.tracer.StartSpan(m.opts.opNameFunc(r), ext.RPCServerOption(ctx))
		ext.HTTPMethod.Set(sp, r.Method)
		ext.HTTPUrl.Set(sp, r.URL.String())
		ext.Component.Set(sp, m.opts.componentName)
		//opts.spanObserver(sp, r)

		r = r.WithContext(opentracing.ContextWithSpan(r.Context(), sp))
		sct := &statusCodeTracker{ResponseWriter: w}
		defer func() {
			ext.HTTPStatusCode.Set(sp, uint16(sct.status))
			if sct.status >= http.StatusInternalServerError || !sct.wroteheader {
				ext.Error.Set(sp, true)
			}
			sp.Finish()
		}()

		h.ServeHTTP(sct.wrappedResponseWriter(), r)
	})
}
