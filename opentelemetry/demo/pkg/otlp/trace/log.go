package trace

import (
	"context"

	"go.opentelemetry.io/otel/api/trace"
	"go.uber.org/zap"
)

// TraceId add trace Id that gets trace id from the context to the logger
func TraceId(ctx context.Context) zap.Field {
	if ctx == nil {
		return zap.Skip()
	}

	spanCtx := trace.RemoteSpanContextFromContext(ctx)
	if !spanCtx.HasTraceID() {
		return zap.Skip()
	}

	traceID := spanCtx.TraceID.String()
	if len(traceID) == 0 {
		// Note the addition of the skip field if
		// a trace is not present in ctx
		return zap.Skip()
	}

	return zap.String("_tid", traceID)
}
