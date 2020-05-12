package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/cenk/backoff"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

// Publisher is a message queue publisher
type Publisher interface {
	Publish(exchange, routingKey string, body []byte) error
}

// NewRabbitMQPublisher is
func NewRabbitMQPublisher(open func(bool) (*amqp.Connection, error), do func(ch *amqp.Channel) error) (Publisher, error) {
	publisher := &rabbitmqPublisher{open: open, do: do, retryTimes: 3, retryInterval: time.Second, logger: logrus.New()}
	if err := publisher.connectWithRetry(); err != nil {
		return nil, err
	}

	return publisher, nil
}

type rabbitmqPublisher struct {
	open    func(bool) (*amqp.Connection, error)
	do      func(ch *amqp.Channel) error
	current *amqp.Connection
	channel *amqp.Channel

	logger        logrus.FieldLogger
	retryTimes    uint64
	retryInterval time.Duration
	retryCh       chan struct{}
	retryMu       sync.Mutex
}

// Publish send a message to the message queue of specified exchange name
func (r *rabbitmqPublisher) Publish(exchange, routingKey string, body []byte) error {
	do := func() error {
		if err := r.publish(exchange, routingKey, body); err != nil {
			if err == amqp.ErrClosed {
				r.logger.WithError(err).Errorln("unavailable network and try to publish the message.")
				r.notifyRetryConnect()
			}
			return err
		}

		return nil
	}

	eboff := backoff.NewExponentialBackOff()
	eboff.InitialInterval = r.retryInterval
	err := backoff.Retry(do, backoff.WithMaxRetries(eboff, r.retryTimes))
	if err != nil {
		return err
	}

	return nil
}

func (r *rabbitmqPublisher) publish(exchange, routingKey string, body []byte) error {
	return r.channel.Publish(
		exchange,   // exchange
		routingKey, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			Body:         body,
		})
}

func (r *rabbitmqPublisher) connect() error {
	var err error
	r.current, err = r.open(false)
	if err != nil {
		return err
	}

	r.channel, err = r.current.Channel()
	if err != nil {
		return err
	}

	if r.do != nil {
		return r.do(r.channel)
	}

	return nil
}

func (r *rabbitmqPublisher) connectWithRetry() error {
	if err := r.connect(); err != nil {
		return err
	}

	r.retryCh = make(chan struct{})
	go func() {
		for {
			select {
			case <-r.retryCh:
				r.retryConnect()
			}
		}
	}()

	return nil
}

func (r *rabbitmqPublisher) notifyRetryConnect() {
	if r.retryCh != nil {
		r.retryCh <- struct{}{}
	}
}

func (r *rabbitmqPublisher) retryConnect() error {
	r.retryMu.Lock()
	defer r.retryMu.Unlock()

	// check if the connect is real closed
	if !r.current.IsClosed() {
		return nil
	}

	// connection already be closed
	current, err := r.open(true)
	if err != nil {
		return err
	}
	r.current = current

	r.channel, err = r.current.Channel()
	return err
}

func NewDefaultRabbitMQStrategy(exchange, bindingKey, qname string) func(ch *amqp.Channel) error {
	return func(ch *amqp.Channel) error {
		err := ch.ExchangeDeclare(
			exchange, // name
			amqp.ExchangeTopic,
			true,  // durable
			false, // auto-deleted
			false, // internal
			false, // no-wait
			nil,   // arguments
		)
		if err != nil {
			return err
		}

		q, err := ch.QueueDeclare(
			qname, // name
			true,  // durable
			false, // delete when unused
			false, // exclusive
			false, // no-wait
			nil,   // arguments
		)
		if err != nil {
			return err
		}

		return ch.QueueBind(
			q.Name,     // queue name
			bindingKey, // routing key
			exchange,   // exchange
			false,
			nil)
	}
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

	strategy := NewDefaultRabbitMQStrategy(exchange, bindingKey, qname)
	r, err := NewRabbitMQPublisher(open, strategy)
	if err != nil {
		panic(err)
	}

	body := []byte("hello")
	go func() {
		for {
			err := r.Publish(exchange, routingKey, body)
			if err != nil {
				fmt.Printf("1. failed to publish the message: %v\n", err)
			} else {
				log.Printf("1. [x] %s Sent %s", routingKey, body)
			}
			//time.Sleep(time.Millisecond * 100)
			time.Sleep(time.Second * 5)
		}
	}()

	go func() {
		for {
			err := r.Publish(exchange, routingKey, body)
			if err != nil {
				fmt.Printf("2. failed to publish the message: %v\n", err)
			} else {
				log.Printf("2. [x] %s Sent %s", routingKey, body)
			}
			//time.Sleep(time.Millisecond * 100)
			time.Sleep(time.Second * 5)
		}
	}()

	select {}
}
