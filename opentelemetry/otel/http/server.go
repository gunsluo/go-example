package main

import (
	"io"
	"log"
	"net/http"

	"go.opentelemetry.io/otel/api/correlation"
	"go.opentelemetry.io/otel/api/global"
	"go.opentelemetry.io/otel/api/trace"
	"go.opentelemetry.io/otel/exporters/trace/stdout"
	"go.opentelemetry.io/otel/instrumentation/httptrace"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func initTracer() {
	// Create stdout exporter to be able to retrieve
	// the collected spans.
	exporter, err := stdout.NewExporter(stdout.Options{PrettyPrint: true})
	if err != nil {
		log.Fatal(err)
	}

	// For the demonstration, use sdktrace.AlwaysSample sampler to sample all traces.
	// In a production application, use sdktrace.ProbabilitySampler with a desired probability.
	tp, err := sdktrace.NewProvider(sdktrace.WithConfig(sdktrace.Config{DefaultSampler: sdktrace.AlwaysSample()}),
		sdktrace.WithSyncer(exporter))
	if err != nil {
		log.Fatal(err)
	}
	global.SetTraceProvider(tp)
}

func main() {
	initTracer()
	tr := global.Tracer("example/server")

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		attrs, entries, spanCtx := httptrace.Extract(req.Context(), req)

		req = req.WithContext(correlation.ContextWithMap(req.Context(), correlation.NewMap(correlation.MapUpdate{
			MultiKV: entries,
		})))

		ctx, span := tr.Start(
			trace.ContextWithRemoteSpanContext(req.Context(), spanCtx),
			"hello",
			trace.WithAttributes(attrs...),
		)
		defer span.End()

		span.AddEvent(ctx, "handling this...")

		_, _ = io.WriteString(w, "Hello, world!\n")
	}

	http.HandleFunc("/hello", helloHandler)
	err := http.ListenAndServe(":7777", nil)
	if err != nil {
		panic(err)
	}
}
