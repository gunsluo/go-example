package main

import (
	"fmt"
	"log"

	"github.com/gunsluo/goews/v2"
)

func main() {
	c, err := goews.NewClient(
		goews.Config{
			Address:  "https://outlook.office365.com/EWS/Exchange.asmx",
			Username: "email@exchangedomain",
			Password: "password",
			Dump:     true,
			NTLM:     false,
			SkipTLS:  false,
		},
	)
	if err != nil {
		log.Fatal("->: ", err.Error())
	}

	err = c.SendEmail(
		"email@exchangedomain",
		[]string{"mhewedy@gmail.com", "someone@else.com"},
		"An email subject",
		"The email body, as plain text",
	)
	if err != nil {
		log.Fatal("err>: ", err.Error())
	}

	fmt.Println("--- success ---")
}
