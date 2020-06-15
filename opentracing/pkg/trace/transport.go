package trace

import (
	"io"
	"net/http"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/log"
)

// Transport wraps a RoundTripper. If a request is being traced with
// Tracer, Transport will inject the current span into the headers,
// and set HTTP related tags on the span.
type Transport struct {
	// The actual RoundTripper to use for the request. A nil
	// RoundTripper defaults to http.DefaultTransport.
	http.RoundTripper

	tr   opentracing.Tracer
	opts *transportOptions
}

type TransportOption func(*transportOptions)

type transportOptions struct {
	componentName            string
	disableClientTrace       bool
	disableInjectSpanContext bool
	operationNameFunc        func(req *http.Request) string
	spanObserver             func(span opentracing.Span, r *http.Request)
}

func TransportOperationNameFunc(operationNameFunc func(req *http.Request) string) TransportOption {
	return func(options *transportOptions) {
		options.operationNameFunc = operationNameFunc
	}
}

func TransportComponentName(componentName string) TransportOption {
	return func(options *transportOptions) {
		options.componentName = componentName
	}
}

func defaultTransportOptions() *transportOptions {
	return &transportOptions{
		componentName: "net/http",
		spanObserver:  func(_ opentracing.Span, _ *http.Request) {},
	}
}

// NewTransport return a transport for http trace
func NewTransport(tr opentracing.Tracer, options ...TransportOption) *Transport {
	transport := &Transport{tr: tr, opts: defaultTransportOptions()}
	for _, opt := range options {
		opt(transport.opts)
	}

	return transport
}

// RoundTrip implements the RoundTripper interface.
func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	rt := t.RoundTripper
	if rt == nil {
		rt = http.DefaultTransport
	}

	sp := t.startSpan(req)
	if sp == nil {
		return rt.RoundTrip(req)
	}

	ext.HTTPMethod.Set(sp, req.Method)
	ext.HTTPUrl.Set(sp, req.URL.String())
	t.opts.spanObserver(sp, req)

	if !t.opts.disableInjectSpanContext {
		carrier := opentracing.HTTPHeadersCarrier(req.Header)
		sp.Tracer().Inject(sp.Context(), opentracing.HTTPHeaders, carrier)
	}

	resp, err := rt.RoundTrip(req)

	if err != nil {
		sp.Finish()
		return resp, err
	}
	ext.HTTPStatusCode.Set(sp, uint16(resp.StatusCode))
	if resp.StatusCode >= http.StatusInternalServerError {
		ext.Error.Set(sp, true)
	}
	if req.Method == "HEAD" {
		sp.Finish()
	} else {
		resp.Body = closeTracker{resp.Body, sp}
	}
	return resp, nil
}

func (t *Transport) startSpan(req *http.Request) opentracing.Span {
	parent := opentracing.SpanFromContext(req.Context())
	var spanctx opentracing.SpanContext
	if parent != nil {
		spanctx = parent.Context()
	}
	var operationName string
	if t.opts.operationNameFunc == nil {
		operationName = req.Method + " " + req.URL.String() + " - HTTP Client"
	} else {
		operationName = t.opts.operationNameFunc(req)
	}
	sp := t.tr.StartSpan(operationName, opentracing.ChildOf(spanctx))

	componentName := t.opts.componentName
	ext.Component.Set(sp, componentName)

	return sp
}

type closeTracker struct {
	io.ReadCloser
	sp opentracing.Span
}

func (c closeTracker) Close() error {
	err := c.ReadCloser.Close()
	c.sp.LogFields(log.String("event", "ClosedBody"))
	c.sp.Finish()
	return err
}
