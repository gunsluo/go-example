package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
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

var mgURL string = "https://api.mailgun.net/v3"

type reader struct {
	*bytes.Reader
}

func newReader(b []byte) *reader {
	return &reader{Reader: bytes.NewReader(b)}
}

func (r *reader) Close() error {
	return nil
}

func main() {
	// Create an instance of the Mailgun Client
	mg := mailgun.NewMailgun(yourDomain, privateAPIKey)
	mg.SetAPIBase(mgURL)

	//sender := "no-reply@luoji.live"
	sender := "ji.luo@target-energysolutions.com"
	//cc := "weilong.yi@target-energysolutions.com"
	//bcc := "xianchao.yu@target-energysolutions.com"
	//sender := "gunsluo@gmail.com"
	subject := "Fancy subject!"
	body := "<p>Hello from Mailgun Go!</p><p>test test</p>"
	recipient := "gunsluo@gmail.com"
	//recipient := "ji.luo@target-energysolutions.com"

	message := mg.NewMessage(sender, subject, body, recipient)
	//message.AddCC(cc)
	//message.AddBCC(bcc)
	message.SetHtml(body)

	fn := "k3.png"
	b, err := ioutil.ReadFile(fn)
	if err != nil {
		panic(err)
	}

	//message.AddAttachment(fn)
	message.AddReaderInline(fn, newReader(b))
	resp, id, err := mg.Send(message)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("ID: %s Resp: %s\n", id, resp)
}
