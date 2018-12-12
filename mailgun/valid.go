package main

import (
	"fmt"
	"log"

	"github.com/mailgun/mailgun-go"
)

// If your plan does not include email validations but you have an account,
// use your public api key (starts with "pubkey-"). If your plan does include
// email validations, use your private api key (starts with "key-")
var apiKey string = ""

func main() {
	// Create an instance of the Validator
	v := mailgun.NewEmailValidator(apiKey)

	verification, err := v.ValidateEmail("no-reply@luoji.live", false)
	//verification, err := v.ValidateEmail("ji.luo@target-energysolutions.com", false)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("result:", verification)
}
