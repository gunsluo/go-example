package main

import (
	"fmt"

	"github.com/sfreiberg/gotwilio"
)

func main() {
	//accountSid := "ABC123..........ABC123"
	//authToken := "ABC123..........ABC123"

	//price()
	prices()
}

func test() {
	accountSid := "AC6cf1019ce309540bd5e0bb99dad92d2d"
	authToken := "c0768fb9ccb116dde9c436704d2b482f"

	twilio := gotwilio.NewTwilioClient(accountSid, authToken)

	from := "+16786715169"
	to := "+8618628076761"
	//from := "+15005550006"
	//to := "+15005550006"
	message := "Welcome to gotwilio!"
	resp, exception, err := twilio.SendSMS(from, to, message, "", "")
	if err != nil {
		panic(err)
	}

	if exception != nil {
		fmt.Printf("exexception-->%+v\n", exception)
	}

	fmt.Printf("resp-->%+v\n", resp)
}

func test1() {
	accountSid := "AC11bec28eb05165e88a141bfd54c78c0e"
	authToken := "7816d68fa7230cef38c124790d5889c1"
	twilio := gotwilio.NewTwilioClient(accountSid, authToken)

	from := "+12053790735"
	to := "+8618583637565"
	message := "Welcome to gotwilio!"
	resp, exception, err := twilio.SendSMS(from, to, message, "", "")
	if err != nil {
		panic(err)
	}

	if exception != nil {
		fmt.Printf("exexception-->%+v\n", exception)
	}

	fmt.Printf("resp-->%+v\n", resp)
}

func price() {
	accountSid := "AC6cf1019ce309540bd5e0bb99dad92d2d"
	authToken := "c0768fb9ccb116dde9c436704d2b482f"
	twilio := gotwilio.NewTwilioClient(accountSid, authToken)

	countryCode := "CN"
	resp, exception, err := twilio.GetSMSPrice(countryCode)
	if err != nil {
		panic(err)
	}

	if exception != nil {
		fmt.Printf("exexception-->%+v\n", exception)
	}

	fmt.Printf("resp-->%+v\n", resp)
}

func prices() {
	accountSid := "AC6cf1019ce309540bd5e0bb99dad92d2d"
	authToken := "c0768fb9ccb116dde9c436704d2b482f"
	twilio := gotwilio.NewTwilioClient(accountSid, authToken)

	resp, exception, err := twilio.GetSMSCountries(
		"",
		//"https://pricing.twilio.com/v1/Messaging/Countries?PageSize=10&Page=1&PageToken=DNAQ",
		&gotwilio.Option{Key: "PageSize", Value: "10"},
		&gotwilio.Option{Key: "Page", Value: "0"},
	)
	if err != nil {
		panic(err)
	}

	if exception != nil {
		fmt.Printf("exexception-->%+v\n", exception)
	}

	fmt.Printf("resp-->%+v\n", resp)
}
