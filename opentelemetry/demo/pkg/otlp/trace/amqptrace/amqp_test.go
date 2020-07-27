package amqptrace

import (
	"context"
	"strings"
	"sync"
	"testing"

	"github.com/streadway/amqp"
	"go.opentelemetry.io/otel/api/trace"
	export "go.opentelemetry.io/otel/sdk/export/trace"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func TestInject(t *testing.T) {
	ctx := context.Background()
	ctx, span := testMockStart(ctx)

	h := amqp.Table{}
	Inject(ctx, h)

	v, ok := h["traceparent"]
	if !ok {
		t.Fatalf("failed to Inject")
	}

	s, ok := v.(string)
	if !ok {
		t.Fatalf("invalid trace id")
	}

	ss := strings.Split(s, "-")
	if len(ss) != 4 {
		t.Fatalf("invalid trace id")
	}

	if span.SpanContext().TraceID.String() != ss[1] {
		t.Fatalf("invalid  trace id")
	}
}

func TestExtract(t *testing.T) {
	ctx := context.Background()
	ctx, span := testMockStart(ctx)

	h := amqp.Table{}
	Inject(ctx, h)

	_, spanCtx := Extract(ctx, h)

	if span.SpanContext().TraceID.String() != spanCtx.TraceID.String() {
		t.Fatalf("failed to Extract")
	}
}

func testMockStart(ctx context.Context) (context.Context, trace.Span) {
	exp := &testExporter{spanMap: make(map[string]*export.SpanData)}
	tp, _ := sdktrace.NewProvider(
		sdktrace.WithSyncer(exp),
		sdktrace.WithConfig(sdktrace.Config{
			DefaultSampler: sdktrace.AlwaysSample(),
		}))

	tracer := tp.Tracer("amqptrace/client")
	ctx, span := tracer.Start(
		ctx,
		"testamqp",
		trace.WithSpanKind(trace.SpanKindClient),
	)
	ctx = trace.ContextWithSpan(ctx, span)

	return ctx, span
}

type testExporter struct {
	mu      sync.Mutex
	spanMap map[string]*export.SpanData
}

func (t *testExporter) ExportSpan(ctx context.Context, s *export.SpanData) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.spanMap[s.Name] = s
}
