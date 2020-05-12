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

// go run receive.go -url amqp://guest:guest@localhost:5671/ -ca-crt=grpc-mtls.dev.meeraspace.com/2_intermediate/certs/ca-chain.cert.pem -client-crt=grpc-mtls.dev.meeraspace.com/4_client/certs/grpc-mtls.dev.meeraspace.com.cert.pem -client-key=grpc-mtls.dev.meeraspace.com/4_client/private/grpc-mtls.dev.meeraspace.com.key.pem
// go run receive.go -url amqps://localhost:5671/ -ca-crt=certs/ca_certificate.pem -client-crt=certs/client/client_certificate.pem -client-key=certs/client/private_key.pem
// go run receive.go -url amqps://guest:guest@localhost:5671/ -ca-crt=certs2/ca.crt -client-crt=certs2/client/client.crt -client-key=certs2/client/client.key
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

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
