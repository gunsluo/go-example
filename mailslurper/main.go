package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"strings"

	"github.com/ory/mail/v3"

	"github.com/emersion/go-smtp"
)

func main() {
	test2()
}

func test() {
	from := "from@example.com"
	recipient := "recipient@example.com"
	// smtps://test:test@mailslurper:1025/?skip_ssl_verify=true
	d := mail.NewDialer("localhost", 2500, "test", "test")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	gm := mail.NewMessage()
	gm.SetHeader("From", from)
	//gm.SetAddressHeader("From", from, fromName)
	gm.SetHeader("To", recipient)
	gm.SetHeader("Subject", "test subject")

	gm.SetBody("text/plain", "test body")
	// Send emails using d.

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
