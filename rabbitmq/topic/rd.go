package main

import (
	"log"
	"os"
	"strings"

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
	var qname string
	if len(os.Args) > 1 {
		qname = os.Args[1]
	}
	log.Println("queue name:", qname)

	var routingKeys = []string{"target"}
	if len(os.Args) > 2 && len(os.Args[2]) > 0 {
		routingKeys = strings.Split(os.Args[2], ",")
	}
	log.Println("routingKeys:", routingKeys)

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		exchange, // name
		"topic",  // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	//exclusive := true
	q, err := ch.QueueDeclare(
		qname, // name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	for _, routingKey := range routingKeys {
		err = ch.QueueBind(
			q.Name,     // queue name
			routingKey, // routing key
			exchange,   // exchange
			false,
			nil)
		failOnError(err, "Failed to bind a queue")
	}
}
