package main

import (
	"fmt"
	"log"

	"github.com/mhewedy/ews"
	"github.com/mhewedy/ews/ewsutil"
)

func main() {

	c := ews.NewClient(
		"https://outlook.office365.com/EWS/Exchange.asmx",
		"email@exchangedomain",
		"password",
		&ews.Config{Dump: true, NTLM: false},
	)

	err := ewsutil.SendEmail(c,
		[]string{"mhewedy@gmail.com", "someone@else.com"},
		"An email subject",
		"The email body, as plain text",
	)

	if err != nil {
		log.Fatal("err>: ", err.Error())
	}

	fmt.Println("--- success ---")
}
