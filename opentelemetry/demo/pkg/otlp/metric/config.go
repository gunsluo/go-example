package metric

import (
	"fmt"
	"os"
	"strconv"
	"time"

	sotlp "github.com/gunsluo/go-example/opentelemetry/demo/pkg/otlp"
	"go.opentelemetry.io/otel/api/metric"
	"go.opentelemetry.io/otel/exporters/otlp"
	"go.opentelemetry.io/otel/sdk/metric/controller/push"
	"go.opentelemetry.io/otel/sdk/metric/processor/basic"
	"go.opentelemetry.io/otel/sdk/metric/selector/simple"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

const (
	// environment variable names
	envEnabled       = "METRIC_ENABLED"
	envAgentEndpoint = "OTLP_AGENT_EDNPOINT"
)

// FromEnv uses environment variables to set the metric's Configuration
func FromEnv() (*Configuration, error) {
	c := &Configuration{}
	return c.FromEnv()
}

// Configuration configures and creates Metric
type Configuration struct {
	// AgentEndpoint is agent address of collecting trace message, host:port
	// Can be provided via environment variable named OTLP_AGENT_EDNPOINT
	AgentEndpoint string

	// Enabled can be provided via environment variable named METRIC_ENABLED
	Enabled bool
}

// FromEnv uses environment variables and overrides existing metric's Configuration
func (c *Configuration) FromEnv() (*Configuration, error) {
	if e := os.Getenv(envEnabled); e != "" {
		if value, err := strconv.ParseBool(e); err == nil {
			c.Enabled = value
		} else {
			return nil, fmt.Errorf("cannot parse env var %s=%s, %w", envEnabled, e, err)
		}
	}

	if e := os.Getenv(envAgentEndpoint); e != "" {
		c.AgentEndpoint = e
	}

	return c, nil
}

// Option is a function that sets some option on the client.
type Option func(c *Options)

// Options control behavior of the client.
type Options struct {
	Logger *zap.Logger
}

// Logger can be provided to logger.
func Logger(logger *zap.Logger) Option {
	return func(c *Options) {
		c.Logger = logger
	}
}

var (
	meterPusher   *push.Controller
	meterProvider metric.Provider
)

// NewMeter create a new metric
func (c *Configuration) NewMeter(name string, options ...Option) (metric.Meter, error) {
	if !c.Enabled {
		return metric.NoopProvider{}.Meter(name), nil
	}

	if c.AgentEndpoint == "" {
		return metric.Meter{}, fmt.Errorf("missing agent address, please set environment variable %s", envAgentEndpoint)
	}

	opts := applyOptions(options...)
	exporter := sotlp.SingletonExporter()
	if exporter == nil {
		exp, err := otlp.NewExporter(otlp.WithInsecure(),
			otlp.WithAddress(c.AgentEndpoint),
			otlp.WithReconnectionPeriod(time.Minute),
			otlp.WithGRPCDialOption(grpc.WithTimeout(5*time.Second)))
		if err != nil {
			return metric.Meter{}, fmt.Errorf("failed to create the collector exporter: %w", err)
		}
		exporter = exp
		sotlp.SetExporter(exporter)
		opts.Logger.With(zap.String("agentEndpoint", c.AgentEndpoint)).Info("success to otlp agent")
	}
	// exporter.Stop()

	if meterPusher == nil {
		meterPusher = push.New(
			basic.New(
				simple.NewWithExactDistribution(),
				exporter,
			),
			exporter,
			push.WithPeriod(30*time.Second),
			//push.WithTimeout(10*time.Second),
		)
		meterProvider = meterPusher.Provider()
		meterPusher.Start()
		opts.Logger.With(zap.String("agentEndpoint", c.AgentEndpoint)).Info("success to create metric pusher and start to push metric")
	}

	return meterProvider.Meter(name), nil
}

func Stop() {
	if meterPusher != nil {
		meterPusher.Stop()
		meterPusher = nil
	}
}

func applyOptions(options ...Option) Options {
	opts := Options{Logger: zap.NewNop()}

	for _, option := range options {
		option(&opts)
	}

	return opts
}
