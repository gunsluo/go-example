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
	var qname string
	if len(os.Args) > 1 && len(os.Args[1]) > 0 {
		qname = os.Args[1]
	}
	log.Println("queue name:", qname)

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	/*
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

				for _, s := range bindingKeys {
			log.Printf("Binding queue %s to exchange %s with routing key %s",
				q.Name, exchange, s)
			err = ch.QueueBind(
				q.Name,   // queue name
				s,        // routing key
				exchange, // exchange
				false,
				nil)
			failOnError(err, "Failed to bind a queue")
		}
	*/

	q, err := ch.QueueDeclare(
		qname, // name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto ack
		false,  // exclusive
		false,  // no local
		false,  // no wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf(" [x] %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}
