package main

import (
	"context"
	"log"

	"go.opentelemetry.io/otel/api/global"
	"go.opentelemetry.io/otel/api/kv"

	"go.opentelemetry.io/otel/exporters/trace/jaeger"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

// initTracer creates a new trace provider instance and registers it as global trace provider.
func initTracer() func() {
	// Create and install Jaeger export pipeline
	_, flush, err := jaeger.NewExportPipeline(
		jaeger.WithCollectorEndpoint("http://localhost:14268/api/traces"),
		jaeger.WithProcess(jaeger.Process{
			ServiceName: "trace-demo",
			Tags: []kv.KeyValue{
				kv.String("exporter", "jaeger"),
				kv.Float64("float", 312.23),
			},
		}),
		jaeger.RegisterAsGlobal(),
		jaeger.WithSDK(&sdktrace.Config{DefaultSampler: sdktrace.AlwaysSample()}),
	)
	if err != nil {
		log.Fatal(err)
	}

	return func() {
		flush()
	}
}

func main() {
	fn := initTracer()
	defer fn()

	ctx := context.Background()

	tr := global.Tracer("component-main")
	//tr.WithSpan(ctx context.Context, spanName string, fn func(ctx context.Context)
	//tr.Start(ctx context.Context, spanName string, opts ...trace.StartOption)
	ctx, span := tr.Start(ctx, "foo")
	//span.SetStatus(codes.Internal, "test-code")
	//span.AddEvent(ctx, "t-event", kv.Key("name").String("luoji"))
	//span.SetAttribute("name", "luoji")
	bar(ctx)
	span.End()
}

func bar(ctx context.Context) {
	tr := global.Tracer("component-bar")
	_, span := tr.Start(ctx, "bar")
	defer span.End()

	// Do bar...
}
