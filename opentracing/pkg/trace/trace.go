package trace

import (
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"github.com/uber/jaeger-client-go/rpcmetrics"
	"github.com/uber/jaeger-lib/metrics"
)

func Init(serviceName string, metricsFactory metrics.Factory, logger logrus.FieldLogger) opentracing.Tracer {
	cfg, err := config.FromEnv()
	if err != nil {
		logger.Fatal("cannot parse Jaeger env vars", err)
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
