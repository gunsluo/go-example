package trace

import (
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"github.com/uber/jaeger-client-go/rpcmetrics"
	"github.com/uber/jaeger-lib/metrics"
	"github.com/uber/jaeger-lib/metrics/expvar"
)

var defaultMetricsFactory = expvar.NewFactory(10)

func Init(serviceName string, logger logrus.FieldLogger, metricsFactory metrics.Factory) opentracing.Tracer {
	cfg, err := config.FromEnv()
	if err != nil {
		logger.Fatal("cannot parse Jaeger env vars", err)
	}

	if metricsFactory == nil {
		metricsFactory = defaultMetricsFactory
	}

	cfg.ServiceName = serviceName
	cfg.Sampler.Type = "const"
	cfg.Sampler.Param = 1

	tracer, _, err := cfg.NewTracer(
		config.Logger(jaeger.StdLogger),
		config.Metrics(metricsFactory),
		config.Observer(rpcmetrics.NewObserver(metricsFactory, rpcmetrics.DefaultNameNormalizer)),
	)
	if err != nil {
		logger.Fatal("cannot initialize Jaeger Tracer", err)
	}

	return tracer
}
