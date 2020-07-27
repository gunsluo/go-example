package amqptrace

import (
	"github.com/streadway/amqp"
	"go.opentelemetry.io/otel/api/global"
	"go.opentelemetry.io/otel/api/propagation"
)

// Option is a function that allows configuration of the amqptrace Extract()
// and Inject() functions
type Option func(*config)

type config struct {
	propagators propagation.Propagators
}

func newConfig(opts []Option) *config {
	c := &config{propagators: global.Propagators()}
	for _, o := range opts {
		o(c)
	}
	return c
}

// WithPropagators sets the propagators to use for Extraction and Injection
func WithPropagators(props propagation.Propagators) Option {
	return func(c *config) {
		c.propagators = props
	}
}

type headerSupplier struct {
	headers amqp.Table
}

func (s *headerSupplier) Get(key string) string {
	value, ok := s.headers[key]
	if !ok {
		return ""
	}

	str, ok := value.(string)
	if !ok {
		return ""
	}

	return str
}

func (s *headerSupplier) Set(key string, value string) {
	s.headers[key] = value
}
