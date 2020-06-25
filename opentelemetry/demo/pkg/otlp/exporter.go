package otlp

import "go.opentelemetry.io/otel/exporters/otlp"

var singleExporter *otlp.Exporter

// SingletonExporter
func SingletonExporter() *otlp.Exporter {
	return singleExporter
}

// SetExporter
func SetExporter(exporter *otlp.Exporter) {
	singleExporter = exporter
}
