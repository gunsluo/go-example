package amqptrace

import (
	"context"

	"github.com/streadway/amqp"
	"go.opentelemetry.io/otel/api/kv"
	"go.opentelemetry.io/otel/api/trace"
)

// Channel is trace channel
type Channel struct {
	Ch     *amqp.Channel
	tracer trace.Tracer
	opts   *channelOptions
}

// ChannelOption is channel option
type ChannelOption func(*channelOptions)

func ChannelComopenentName(componentName string) ChannelOption {
	return func(o *channelOptions) {
		o.componentName = componentName
	}
}

type channelOptions struct {
	componentName string
}

func applyChannelOptions(opts ...ChannelOption) *channelOptions {
	o := &channelOptions{componentName: "net/amqp"}

	for _, opt := range opts {
		opt(o)
	}
	return o
}

// NewChannel create a trace channel
func NewChannel(tracer trace.Tracer, channel *amqp.Channel, opts ...ChannelOption) *Channel {
	o := applyChannelOptions(opts...)
	ch := &Channel{tracer: tracer, Ch: channel, opts: o}

	return ch
}

// Publish sends a Publishing from the client to an exchange on the server.
func (ch *Channel) Publish(ctx context.Context, exchange, key string, mandatory, immediate bool, msg amqp.Publishing) error {
	var span trace.Span
	ctx, span = ch.tracer.Start(
		ctx,
		exchange+"."+key+" - Rabbitmq Publisher",
		trace.WithSpanKind(trace.SpanKindClient),
		trace.WithAttributes(
			kv.String("exchange", exchange),
			kv.String("routekey", key),
			kv.String("content-type", msg.ContentType),
			kv.Key("component").String(ch.opts.componentName),
		),
	)
	defer span.End()

	if msg.Headers == nil {
		msg.Headers = amqp.Table{}
	}
	Inject(ctx, msg.Headers)

	if err := ch.Ch.Publish(exchange, key, mandatory, immediate, msg); err != nil {
		span.SetAttribute("error", err)
		return err
	}

	return nil
}
