package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.opentelemetry.io/otel/api/global"
	"go.opentelemetry.io/otel/api/kv"
	"go.opentelemetry.io/otel/api/metric"
	"go.opentelemetry.io/otel/exporters/otlp"
	"go.opentelemetry.io/otel/sdk/metric/controller/push"
	"go.opentelemetry.io/otel/sdk/metric/selector/simple"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

// Initializes an OTLP exporter, and configures the corresponding trace and
// metric providers.
func initProvider() (*otlp.Exporter, *push.Controller) {
	exp, err := otlp.NewExporter(
		otlp.WithInsecure(),
		otlp.WithAddress("localhost:55680"),
	)
	handleErr(err, "Failed to create exporter: $v")

	traceProvider, err := sdktrace.NewProvider(
		sdktrace.WithConfig(sdktrace.Config{DefaultSampler: sdktrace.AlwaysSample()}),
		sdktrace.WithResource(resource.New(
			// the service name used to display traces in Jaeger
			kv.Key("service.name").String("msg-service"),
		)),
		sdktrace.WithSyncer(exp),
	)
	handleErr(err, "Failed to create trace provider: %v")

	pusher := push.New(
		simple.NewWithExactDistribution(),
		exp,
		push.WithStateful(true),
		push.WithPeriod(2*time.Second),
	)

	global.SetTraceProvider(traceProvider)
	global.SetMeterProvider(pusher.Provider())
	pusher.Start()

	return exp, pusher
}

func main() {
	exp, pusher := initProvider()
	defer func() { handleErr(exp.Stop(), "Failed to stop exporter") }()
	defer pusher.Stop() // pushes any last exports to the receiver

	tracer := global.Tracer("mage-sense")
	meter := global.Meter("mage-read")

	// labels represent additional descriptors that can be bound to a metric
	// observer or recorder. In this case they describe the location in
	// which a spell scroll is scribed.
	commonLabels := []kv.KeyValue{
		kv.String("work-room", "East Scriptorium"),
		kv.String("occupancy", "69,105"),
		kv.String("priority", "Ultra"),
	}

	// Observer metric example
	oneMetricCB := func(_ context.Context, result metric.Float64ObserverResult) {
		result.Observe(1, commonLabels...)
	}
	_ = metric.Must(meter).NewFloat64ValueObserver("scrying.glass.one", oneMetricCB,
		metric.WithDescription("A ValueObserver set to 1.0"),
	)

	// Recorder metric example
	valuerecorder := metric.Must(meter).
		NewFloat64ValueRecorder("scrying.glass.two").
		Bind(commonLabels...)
	defer valuerecorder.Unbind()

	// work begins
	ctx, span := tracer.Start(context.Background(), "Archmage-Overlord-Inspection")
	for i := 0; i < 10; i++ {
		_, innerSpan := tracer.Start(ctx, fmt.Sprintf("Minion-%d", i))
		log.Println("Minions hard at work, scribing...")
		valuerecorder.Record(ctx, float64(i)*1.5)

		<-time.After(time.Second)
		innerSpan.End()
	}

	span.End()
	<-time.After(time.Second)

	log.Println("Spell-scroll scribed!")
}

func handleErr(err error, message string) {
	if err != nil {
		log.Fatalf("%s: %v", message, err)
	}
}
