package record

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/gunsluo/go-example/opentelemetry/demo/pkg/otlp/trace"
	"github.com/gunsluo/go-example/opentelemetry/demo/pkg/otlp/trace/amqptrace"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

// Server implements jaeger-demo-frontend service
type Server struct {
	logger *zap.Logger

	rabbitmqUrl  string
	rabbitmqConn *amqp.Connection
}

// ConfigOptions used to make sure service clients
// can find correct server ports
type ConfigOptions struct {
	MQUrl    string
	MQPrefix string
}

// NewServer creates a new frontend.Server
func NewServer(options ConfigOptions, logger *zap.Logger) (*Server, error) {
	logger = logger.Named("record")
	s := &Server{
		rabbitmqUrl: options.MQUrl,
		logger:      logger,
	}

	// rabbitmq
	rabbitmqConn, err := amqp.Dial(options.MQUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to rabbitmq, %s: %w", options.MQUrl, err)
	}
	s.rabbitmqConn = rabbitmqConn
	s.logger.With(zap.String("mq-url", options.MQUrl)).Info("success to connect to rabbitmq")

	return s, nil
}

// Run starts the frontend server
func (s *Server) Run() error {
	rabbitmqCh, err := s.rabbitmqConn.Channel()
	if err != nil {
		return fmt.Errorf("failed to create to channel, %w", err)
	}

	queue, err := rabbitmqCh.QueueDeclare(
		"test-record", // name
		false,         // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	if err != nil {
		return fmt.Errorf("failed to create to queue, %w", err)
	}

	delivery, err := rabbitmqCh.Consume(
		queue.Name, // queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	if err != nil {
		return fmt.Errorf("failed to create to consume, %w", err)
	}

	// trace
	traceConfig, err := trace.FromEnv()
	if err != nil {
		return fmt.Errorf("failed to loading trace config, %w", err)
	}
	traceConfig.ServiceName = "record"

	tracer, err := traceConfig.NewTracer(trace.ServiceName("record"), trace.Logger(s.logger))
	if err != nil {
		return fmt.Errorf("failed to create trace, %w", err)
	}

	s.logger.With(zap.String("mq-url", s.rabbitmqUrl)).Info("starting up record server")
	consumer := amqptrace.NewConsumer(tracer, delivery, s.handle, amqptrace.ConsumeComopenentName("Record Rabbimq Server"))
	consumer.Accpet()

	return nil
}

type user struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	CertId string `json:"cert_id"`
}

func (s *Server) handle(ctx context.Context, msg amqp.Delivery) error {
	u := user{}
	if err := json.Unmarshal(msg.Body, &u); err != nil {
		log.Printf("failed to received a message: %v", err)
		return err
	}

	s.logger.With(trace.TraceId(ctx), zap.Any("user", u)).Info("Received a record message")
	return nil
}
