module github.com/gunsluo/go-example/opentelemetry/demo

go 1.14

require (
	github.com/gogo/protobuf v1.3.1
	github.com/jmoiron/sqlx v1.2.1-0.20190826204134-d7d95172beb5
	github.com/lib/pq v1.7.1
	github.com/luna-duclos/instrumentedsql v1.1.3
	github.com/rs/cors v1.7.0
	github.com/spf13/cobra v1.0.0
	github.com/streadway/amqp v0.0.0-20190827072141-edfb9018d271
	github.com/xo/dburl v0.0.0-20200124232849-e9ec94f52bc3
	go.opentelemetry.io/collector v0.9.0
	go.opentelemetry.io/contrib/instrumentation/net/http/httptrace v0.11.0
	go.opentelemetry.io/otel v0.11.0
	go.opentelemetry.io/otel/exporters/otlp v0.11.0
	go.opentelemetry.io/otel/sdk v0.11.0
	go.uber.org/zap v1.15.0
	google.golang.org/grpc v1.31.0
)
