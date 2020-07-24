package record

import (
	"encoding/json"
	"fmt"
	"log"

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

	s.logger.With(zap.String("mq-url", s.rabbitmqUrl)).Info("starting up record server")
	for msg := range delivery {
		go s.handle(msg)
	}

	return nil
}

type user struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	CertId string `json:"cert_id"`
}

func (s *Server) handle(msg amqp.Delivery) {
	u := user{}
	if err := json.Unmarshal(msg.Body, &u); err != nil {
		log.Printf("failed to received a message: %v", err)
		return
	}

	s.logger.With(zap.Any("user", u)).Info("Received a record message")
}
