package foo

import (
	"context"

	"go.opentelemetry.io/otel/api/global"
	"go.opentelemetry.io/otel/api/kv"
	"go.opentelemetry.io/otel/api/trace"
)

var (
	lemonsKey = kv.Key("ex.com/lemons")
)

// SubOperation is an example to demonstrate the use of named tracer.
// It creates a named tracer with its package path.
func SubOperation(ctx context.Context) error {

	// Using global provider. Alternative is to have application provide a getter
	// for its component to get the instance of the provider.
	tr := global.Tracer("example/namedtracer/foo")
	return tr.WithSpan(
		ctx,
		"Sub operation...",
		func(ctx context.Context) error {
			trace.SpanFromContext(ctx).SetAttributes(lemonsKey.String("five"))

			trace.SpanFromContext(ctx).AddEvent(ctx, "Sub span event")

			return nil
		},
	)
}
