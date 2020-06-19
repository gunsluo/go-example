package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"go.opentelemetry.io/otel/api/kv"

	"net/http"
	"time"

	"go.opentelemetry.io/otel/api/correlation"
	"go.opentelemetry.io/otel/api/global"
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
	url := flag.String("server", "http://localhost:7777/hello", "server url")
	flag.Parse()

	client := http.DefaultClient
	ctx := correlation.NewContext(context.Background(),
		kv.String("username", "donuts"),
	)

	var body []byte

	tr := global.Tracer("example/client")
	err := tr.WithSpan(ctx, "say hello",
		func(ctx context.Context) error {
			req, _ := http.NewRequest("GET", *url, nil)

			ctx, req = httptrace.W3C(ctx, req)
			httptrace.Inject(ctx, req)

			fmt.Printf("Sending request...\n")
			res, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			body, err = ioutil.ReadAll(res.Body)
			_ = res.Body.Close()

			return err
		})

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response Received: %s\n\n\n", body)
	fmt.Printf("Waiting for few seconds to export spans ...\n\n")
	time.Sleep(10 * time.Second)
	fmt.Printf("Inspect traces on stdout\n")
}
