package main

import (
	"context"
	"log"

	"github.com/gunsluo/go-example/opentelemetry/otel/namedtracer/foo"
	"go.opentelemetry.io/otel/api/correlation"
	"go.opentelemetry.io/otel/api/global"
	"go.opentelemetry.io/otel/api/kv"
	"go.opentelemetry.io/otel/api/trace"
	"go.opentelemetry.io/otel/exporters/trace/stdout"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

var (
	fooKey     = kv.Key("ex.com/foo")
	barKey     = kv.Key("ex.com/bar")
	anotherKey = kv.Key("ex.com/another")
)

var tp *sdktrace.Provider

// initTracer creates and registers trace provider instance.
func initTracer() {
	var err error
	exp, err := stdout.NewExporter(stdout.Options{})
	if err != nil {
		log.Panicf("failed to initialize stdout exporter %v\n", err)
		return
	}
	tp, err = sdktrace.NewProvider(sdktrace.WithSyncer(exp),
		sdktrace.WithConfig(sdktrace.Config{DefaultSampler: sdktrace.AlwaysSample()}))
	if err != nil {
		log.Panicf("failed to initialize trace provider %v\n", err)
	}
	global.SetTraceProvider(tp)
}

func main() {
	// initialize trace provider.
	initTracer()

	// Create a named tracer with package path as its name.
	tracer := tp.Tracer("example/namedtracer/main")
	ctx := context.Background()

	ctx = correlation.NewContext(ctx,
		fooKey.String("foo1"),
		barKey.String("bar1"),
	)

	err := tracer.WithSpan(ctx, "operation", func(ctx context.Context) error {

		trace.SpanFromContext(ctx).AddEvent(ctx, "Nice operation!", kv.Key("bogons").Int(100))

		trace.SpanFromContext(ctx).SetAttributes(anotherKey.String("yes"))

		return foo.SubOperation(ctx)
	})
	if err != nil {
		panic(err)
	}
}
