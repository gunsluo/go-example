package amqptrace

import (
	"context"

	"github.com/streadway/amqp"
	"go.opentelemetry.io/otel/api/correlation"
	"go.opentelemetry.io/otel/api/kv"
	"go.opentelemetry.io/otel/api/trace"
)

// Consumer is rabbit consumer for trace
type Consumer struct {
	delivery <-chan amqp.Delivery
	tracer   trace.Tracer
	handle   func(context.Context, amqp.Delivery) error
	opts     *consumeOptions
}

// ConsumeOption is channel option
type ConsumeOption func(*consumeOptions)

func ConsumeComopenentName(componentName string) ConsumeOption {
	return func(o *consumeOptions) {
		o.componentName = componentName
	}
}

type consumeOptions struct {
	componentName string
}

func applyConsumeOptions(opts ...ConsumeOption) *consumeOptions {
	o := &consumeOptions{componentName: "net/amqp"}
	for _, opt := range opts {
		opt(o)
	}

	return o
}

// NewConsumer create a trace consumer
func NewConsumer(tracer trace.Tracer, delivery <-chan amqp.Delivery, handle func(context.Context, amqp.Delivery) error, opts ...ConsumeOption) *Consumer {
	o := applyConsumeOptions(opts...)
	c := &Consumer{tracer: tracer, delivery: delivery, handle: handle, opts: o}

	return c
}

// Accept accept a message from rabbitmq
func (c *Consumer) Accpet() {
	for msg := range c.delivery {
		go c.accept(msg)
	}
}

func (c *Consumer) accept(msg amqp.Delivery) {
	ctx := context.Background()
	entries, spCtx := Extract(ctx, msg.Headers)
	ctx = correlation.ContextWithMap(ctx, correlation.NewMap(correlation.MapUpdate{
		MultiKV: entries,
	}))

	ctx, span := c.tracer.Start(
		trace.ContextWithRemoteSpanContext(ctx, spCtx),
		msg.Exchange+"."+msg.RoutingKey,
		trace.WithSpanKind(trace.SpanKindServer),
		trace.WithAttributes(
			kv.String("exchange", msg.Exchange),
			kv.String("routekey", msg.RoutingKey),
			kv.String("content-type", msg.ContentType),
			kv.Key("component").String(c.opts.componentName),
		),
	)
	defer span.End()

	if err := c.handle(ctx, msg); err != nil {
		span.SetAttribute("error", err)
	}
}
