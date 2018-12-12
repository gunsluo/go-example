package main

import (
	"fmt"
	"log"

	"github.com/mailgun/mailgun-go"
)

// Your available domain names can be found here:
// (https://app.mailgun.com/app/domains)
var yourDomain string = "luoji.live" // e.g. mg.yourcompany.com

// The API Keys are found in your Account Menu, under "Settings":
// (https://app.mailgun.com/app/account/security)

// starts with "key-"
var privateAPIKey string = ""

func main() {
	// Create an instance of the Mailgun Client
	mg := mailgun.NewMailgun(yourDomain, privateAPIKey)

	sender := "no-reply@luoji.live"
	subject := "Fancy subject!"
	body := "Hello from Mailgun Go!"
	recipient := "ji.luo@target-energysolutions.com"

	sendMessage(mg, sender, subject, body, recipient)
}

func sendMessage(mg mailgun.Mailgun, sender, subject, body, recipient string) {
	message := mg.NewMessage(sender, subject, body, recipient)
	resp, id, err := mg.Send(message)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("ID: %s Resp: %s\n", id, resp)
}
