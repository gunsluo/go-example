package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

type sender2Receiver struct {
	addr      string
	exchange  string
	queueName string

	conn *amqp.Connection

	// channel of the producter
	sch *amqp.Channel

	// channel of the consumer
	cch      *amqp.Channel
	queue    amqp.Queue
	delivery <-chan amqp.Delivery
}

func newSender2Receiver(rabbitmqAddr string) (*sender2Receiver, error) {
	s := &sender2Receiver{addr: rabbitmqAddr, exchange: "bus.ex", queueName: ""}
	s.initialize()

	return s, nil
}

func (s *sender2Receiver) initialize() error {
	// initialize amqp connection
	conn, err := amqp.Dial(s.addr)
	if err != nil {
		return err
	}
	s.conn = conn

	// initialize the channel of the producter
	sch, err := conn.Channel()
	if err != nil {
		return err
	}

	// declare exchange on channel
	err = sch.ExchangeDeclare(
		s.exchange, // name
		"fanout",   // type
		true,       // durable
		false,      // auto-deleted
		false,      // internal
		false,      // no-wait
		nil,        // arguments
	)
	if err != nil {
		return err
	}
	s.sch = sch

	// initialize the channel of the consumer
	cch, err := conn.Channel()
	if err != nil {
		return err
	}
	s.cch = cch

	q, err := cch.QueueDeclare(
		s.queueName, // name
		false,       // durable
		false,       // delete when usused
		true,        // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	if err != nil {
		return err
	}
	s.queue = q

	err = cch.QueueBind(
		q.Name,     // queue name
		"",         // routing key
		s.exchange, // exchange
		false,
		nil)
	if err != nil {
		return err
	}

	delivery, err := s.cch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	s.delivery = delivery

	return nil
}

func (s *sender2Receiver) send(msg []byte) error {
	err := s.sch.Publish(
		s.exchange, // exchange
		"",         // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        msg,
		})

	return err
}

func (s *sender2Receiver) accept() []byte {
	select {
	case d := <-s.delivery:
		d.Ack(false)
		return d.Body
	}
}

func (s *sender2Receiver) run() {
	for {
		msg := s.accept()
		fmt.Println("receive:", string(msg))
	}
}

func main() {
	ss, err := newSender2Receiver("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	go ss.run()

	ss.send([]byte("hello"))
	ss.send([]byte("world"))

	select {}
}
