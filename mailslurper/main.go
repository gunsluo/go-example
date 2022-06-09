package main

import (
	"context"
	"crypto/tls"
	"fmt"

	"github.com/ory/mail/v3"
)

func main() {
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
