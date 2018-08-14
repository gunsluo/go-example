package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/streadway/amqp"
	"gitlab.com/target-digital-transformation/ses/mq"
)

const (
	sesExchange = "ses.en"
)

func RoutingKey(keys ...string) string {
	if len(keys) == 0 {
		return ""
	}

	return fmt.Sprintf("ses.%s", strings.Join(keys, "."))
}

func BindingKey(key string) string {
	return fmt.Sprintf("ses.#.%s.#", key)
}

func QueueName(key string) string {
	return fmt.Sprintf("ses-%s.qn", key)
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	routingKeys := os.Args[1:]
	if len(routingKeys) == 0 {
		routingKeys = append(routingKeys, "smtp")
	}

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	options := &mq.RabbitMQOptions{}
	options.ExchangeDeclare(sesExchange, amqp.ExchangeTopic).
		QueueDeclare(QueueName("smtp"), true).
		QueueDeclare(QueueName("aws"), true).
		QueueDeclare(QueueName("mailgun"), true).
		Bind(sesExchange, QueueName("smtp"), []string{BindingKey("smtp")}).
		Bind(sesExchange, QueueName("aws"), []string{BindingKey("aws")}).
		Bind(sesExchange, QueueName("mailgun"), []string{BindingKey("mailgun")})

	publisher, err := mq.NewRabbitMQPublisher(conn, options)
	failOnError(err, "Failed to create publisher")

	body := []byte("hello")
	err = publisher.Publish(sesExchange, RoutingKey(routingKeys...), body)
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", body)
}
