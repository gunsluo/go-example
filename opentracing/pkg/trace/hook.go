package trace

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"github.com/uber/jaeger-client-go"
)

type InjectHook struct {
	InjectFunc func(context.Context) logrus.Fields
}

func (hook *InjectHook) Fire(entry *logrus.Entry) error {
	if hook.InjectFunc == nil {
		return nil
	}

	fields := hook.InjectFunc(entry.Context)
	for k, v := range fields {
		entry.Data[k] = v
	}

	return nil
}

func (hook *InjectHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func DefaultInjectHookFunc(ctx context.Context) logrus.Fields {
	if ctx == nil {
		return nil
	}

	sp := opentracing.SpanFromContext(ctx)
	if sp == nil {
		return nil
	}

	spCtx, ok := sp.Context().(jaeger.SpanContext)
	if !ok {
		return nil
	}

	return logrus.Fields{"_tid": spCtx.TraceID().String()}
}
