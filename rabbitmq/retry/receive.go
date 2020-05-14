package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

// Consumer is a message queue consumer
type Consumer interface {
	Deliver(do func(amqp.Delivery)) Consumer
	Begin() error
	End() error
}

type rabbitmqConsumer struct {
	open     func(bool) (*amqp.Connection, error)
	current  *amqp.Connection
	channel  *amqp.Channel
	delivery <-chan amqp.Delivery

	logger logrus.FieldLogger
	qname  string
	do     func(amqp.Delivery)

	shutdown    chan struct{}
	maxInterval time.Duration
}

// Deliver set a deliver func
func (r *rabbitmqConsumer) Deliver(do func(amqp.Delivery)) *rabbitmqConsumer {
	r.do = do
	return r
}

func (r *rabbitmqConsumer) Begin() error {
	if err := r.connect(); err != nil {
		return err
	}
	r.shutdown = make(chan struct{})
	r.maxInterval = time.Minute

	delivery, err := r.channel.Consume(
		r.qname, // queue
		"",      // consumer
		false,   // auto ack
		false,   // exclusive
		false,   // no local
		false,   // no wait
		nil,     // args
	)
	if err != nil {
		return err
	}
	r.delivery = delivery

	r.logger.Infoln("begin to accept messages")
	go func() {
		for {
			select {
			case <-r.shutdown:
				break
			case msg, ok := <-r.delivery:
				if !ok {
					// try to connect ...
					r.waitRepairConnect()
					continue
				}

				r.do(msg)
			}
		}

		r.logger.Infoln("stop to accept messages")
	}()

	return nil
}

func (r *rabbitmqConsumer) waitRepairConnect() {
	retries := 1
	for {
		r.logger.Infof("check network and then accept messages, try %d times", retries)
		if err := r.repairConnect(); err == nil {
			break
		}

		retries++
		d := time.Duration(5+rand.Intn(5)+2*retries) * time.Second
		if d > r.maxInterval {
			d = r.maxInterval
		}
		time.Sleep(d)
	}
}

func (r *rabbitmqConsumer) repairConnect() error {
	if r.current.IsClosed() {
		current, err := r.open(true)
		if err != nil {
			return err
		}
		r.current = current

		channel, err := r.current.Channel()
		if err != nil {
			return err
		}
		r.channel = channel

		delivery, err := r.channel.Consume(
			r.qname, // queue
			"",      // consumer
			false,   // auto ack
			false,   // exclusive
			false,   // no local
			false,   // no wait
			nil,     // args
		)
		if err != nil {
			return err
		}
		r.delivery = delivery
	}

	return nil
}

func (r *rabbitmqConsumer) connect() error {
	var err error
	r.current, err = r.open(false)
	if err != nil {
		return err
	}

	r.channel, err = r.current.Channel()
	return err
}

func (r *rabbitmqConsumer) End() error {
	close(r.shutdown)
	return nil
}

var (
	exchange   = "ses.agent.e"
	qname      = "ses.agent.q"
	routingKey = "ses.agent.aws"
	bindingKey = "ses.agent.aws"
)

func main() {
	var dsn = "amqp://guest:guest@localhost:5672/"

	var conn *amqp.Connection
	open := func(force bool) (*amqp.Connection, error) {
		if !force && conn != nil {
			return conn, nil
		}

		var err error
		conn, err = amqp.Dial(dsn)
		return conn, err
	}

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	c := rabbitmqConsumer{open: open, logger: logrus.New(), qname: qname}
	err := c.Deliver(func(msg amqp.Delivery) {
		log.Printf(" [x] %s", msg.Body)
		msg.Ack(false)
	}).Begin()
	if err != nil {
		panic(err)
	}

	select {}
}
