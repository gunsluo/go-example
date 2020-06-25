package trace

import (
	"go.opentelemetry.io/otel/api/trace"
)

// FromEnv uses environment variables to set the tracer's Configuration
func FromEnv() (*Configuration, error) {
	c := &Configuration{}
	return c.FromEnv()
}

// NewTracerFromEnv uses environment variables to create the tracer
func NewTracerFromEnv(options ...Option) (trace.Tracer, error) {
	c, err := FromEnv()
	if err != nil {
		return nil, err
	}

	return c.NewTracer(options...)
}
