package main

import (
	"log"
	"os"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

const (
	exchange = "sync"
)

func main() {
	var sendRoutingKey = "target"
	if len(os.Args) > 1 && len(os.Args[1]) > 0 {
		sendRoutingKey = os.Args[1]
	}
	log.Println("sendRoutingKey:", sendRoutingKey)

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	body := "hello"
	err = ch.Publish(
		exchange,       // exchange
		sendRoutingKey, // routing key
		false,          // mandatory
		false,          // immediate
		amqp.Publishing{
			ContentType:  "text/plain",
			Body:         []byte(body),
			DeliveryMode: amqp.Persistent,
		})
	failOnError(err, "Failed to publish a message")

	log.Printf(" [x] %s Sent %s", sendRoutingKey, body)
}
