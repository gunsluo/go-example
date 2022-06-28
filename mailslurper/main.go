package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"strings"
	"time"

	"github.com/ory/mail/v3"

	"github.com/emersion/go-smtp"
)

func main() {
	test()
}

func test() {
	//from := "from@example.com"
	//recipient := "recipient@example.com"
	from := "no-reply@ory.kratos.sh"
	recipient := "recipient@example.com"
	// smtps://test:test@mailslurper:1025/?skip_ssl_verify=true
	/*
		d := mail.NewDialer("mailslurper.infra", 2500, "test", "test")
		//d := mail.NewDialer("127.0.0.1", 2500, "test", "test")
		d.TLSConfig = &tls.Config{InsecureSkipVerify: false}
	*/
	d := &mail.Dialer{
		Host:         "mailslurper.infra",
		Port:         2500,
		Username:     "test",
		Password:     "test",
		LocalName:    "localhost",
		Timeout:      time.Second * 10,
		RetryFailure: true,
	}

	d.TLSConfig = &tls.Config{InsecureSkipVerify: true, Certificates: []tls.Certificate{}, ServerName: "mailslurper.infra"}
	// Enforcing StartTLS
	//d.StartTLSPolicy = mail.MandatoryStartTLS
	//fmt.Println("2--->", skipStartTLS)

	//dialer := &gomail.Dialer{}

	gm := mail.NewMessage()
	gm.SetHeader("From", from)
	//gm.SetAddressHeader("From", from, fromName)
	gm.SetHeader("To", recipient)
	gm.SetHeader("Subject", "test subject")

	gm.SetBody("text/plain", "test body")

	htmlBody := `Hi, please verify your account by clicking the following link:

<a href="http://127.0.0.1:4433/self-service/verification?flow=65622daa-b9a8-48f4-a838-f3b1c91df6ba&amp;token=3rLxoXICxk9nHGG0G1uAqzrZkKp1gIuA">http://127.0.0.1:4433/self-service/verification?flow=65622daa-b9a8-48f4-a838-f3b1c91df6ba&amp;token=3rLxoXICxk9nHGG0G1uAqzrZkKp1gIuA</a>`
	gm.AddAlternative("text/html", htmlBody)
	// Send emails using d.

	//message_from=no-reply@ory.kratos.sh service_name=Ory Kratos service_version=master smtp_server=mailslurper.infra:2500 smtp_ssl_enabled=false
	fmt.Println("--->", d.SSL)
	ctx := context.Background()
	if err := d.DialAndSend(ctx, gm); err != nil {
		panic(err)
	}

	fmt.Println("success to send")
}

func test2() {
	from := "from@example.com"
	recipient := "recipient@example.com"

	// Setup authentication information.
	//auth := sasl.NewPlainClient("", "test", "test")

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	to := []string{recipient}
	msg := strings.NewReader("To: recipient@example.net\r\n" +
		"Subject: discount Gophers!\r\n" +
		"\r\n" +
		"This is the email body.\r\n")
	err := smtp.SendMail("localhost:2500", nil, from, to, msg)
	if err != nil {
		panic(err)
	}

	fmt.Println("success to send")
}
