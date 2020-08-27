package amqptrace

import (
	"context"

	"github.com/streadway/amqp"
	"go.opentelemetry.io/otel/api/correlation"
	"go.opentelemetry.io/otel/api/propagation"
	"go.opentelemetry.io/otel/api/trace"
	"go.opentelemetry.io/otel/label"
)

// Inject injects correlation context and span context into the gRPC
// metadata object. This function is meant to be used on outgoing
// requests.
func Inject(ctx context.Context, headers amqp.Table, opts ...Option) {
	c := newConfig(opts)
	propagation.InjectHTTP(ctx, c.propagators, &headerSupplier{
		headers: headers,
	})
}

// Extract returns the correlation context and span context that
// another service encoded in the rabbitmq Table object with Inject.
// This function is meant to be used on incoming requests.
func Extract(ctx context.Context, headers amqp.Table, opts ...Option) ([]label.KeyValue, trace.SpanContext) {
	c := newConfig(opts)
	ctx = propagation.ExtractHTTP(ctx, c.propagators, &headerSupplier{
		headers: headers,
	})

	spanContext := trace.RemoteSpanContextFromContext(ctx)
	var correlationCtxKVs []label.KeyValue
	correlation.MapFromContext(ctx).Foreach(func(kv label.KeyValue) bool {
		correlationCtxKVs = append(correlationCtxKVs, kv)
		return true
	})

	return correlationCtxKVs, spanContext
}
