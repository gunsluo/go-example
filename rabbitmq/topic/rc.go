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

	var sendRoutingKey = "target"
	if len(os.Args) > 3 && len(os.Args[3]) > 0 {
		sendRoutingKey = os.Args[3]
	}
	log.Println("sendRoutingKey:", sendRoutingKey)

	/*
		var routingKey = "target"
		if len(os.Args) > 1 {
			routingKey = os.Args[1]
		}
		log.Println("routingKey:", routingKey)

		var qname string
		if len(os.Args) > 2 && len(os.Args[2]) > 0 {
			qname = os.Args[2]
		}
		log.Println("queue name:", qname)
	*/

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

	if qname != "" {
		//exclusive := true
		q, err := ch.QueueDeclare(
			qname, // name
			false, // durable
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

	body := "hello"
	err = ch.Publish(
		exchange,       // exchange
		sendRoutingKey, // routing key
		false,          // mandatory
		false,          // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")

	log.Printf(" [x] %s Sent %s", sendRoutingKey, body)
}
