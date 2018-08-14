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
	if len(os.Args) < 2 {
		log.Printf("Usage: %s [binding_key]...", os.Args[0])
		os.Exit(0)
	}

	agentKey := os.Args[1]

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

	consumer, err := mq.NewRabbitMQConsumer(conn, options)
	failOnError(err, "Failed to create consumer")

	msgs, err := consumer.Consume(QueueName(agentKey))
	failOnError(err, "Failed to consume a message")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf(" [x] %s", d.Body)
			d.Ack(false)
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}
