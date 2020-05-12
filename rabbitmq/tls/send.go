package main

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func newConn(url, caPath, clientCrt, clientKey string) (*amqp.Connection, error) {
	if clientCrt == "" && caPath == "" && clientKey == "" {
		return amqp.Dial(url)
	}

	cPool := x509.NewCertPool()
	if caPath != "" {
		caCert, err := ioutil.ReadFile(caPath)
		if err != nil {
			return nil, fmt.Errorf("invalid CA crt file: %s", caPath)
		}
		if cPool.AppendCertsFromPEM(caCert) != true {
			panic(errors.New("failed to append CA crt"))
		}
	}

	if clientCrt == "" && clientKey == "" {
		clientTLSConfig := &tls.Config{
			//InsecureSkipVerify: true,
			RootCAs: cPool,
		}
		return amqp.DialTLS(url, clientTLSConfig)
	}

	clientCert, err := tls.LoadX509KeyPair(clientCrt, clientKey)
	if err != nil {
		return nil, fmt.Errorf("invalid client crt file: %s %s", clientCrt, clientKey)
	}

	clientTLSConfig := &tls.Config{
		//InsecureSkipVerify: true,
		RootCAs:      cPool,
		Certificates: []tls.Certificate{clientCert},
		//ServerName:   *serverName,
	}

	fmt.Println("set certificate")
	return amqp.DialTLS(url, clientTLSConfig)
}

// go run send.go -url amqps://localhost:5671/ -ca-crt=certs/ca_certificate.pem -client-crt=certs/client/client_certificate.pem -client-key=certs/client/private_key.pem
func main() {
	url := flag.String("url", "amqp://guest:guest@localhost:5672/", "rabbitmq url")
	caCrt := flag.String("ca-crt", "", "CA certificate")
	clientCrt := flag.String("client-crt", "", "Client certificate")
	clientKey := flag.String("client-key", "", "Client key")
	flag.Parse()

	conn, err := newConn(*url, *caCrt, *clientCrt, *clientKey)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	body := "hello"
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	log.Printf(" [x] Sent %s", body)
	failOnError(err, "Failed to publish a message")
}
