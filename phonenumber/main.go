package main

import (
	"fmt"

	"github.com/nyaruka/phonenumbers"
)

func main() {
	//number := "+96819999999999"
	number := "+92 04658863211"
	//number := "+86 18980501737"
	//number := "+96819960515"
	//number := "6502530000"
	// 6502530000
	// parse our phone number

	//num, err := phonenumbers.ParseAndKeepRawInput(number, "")
	num, err := phonenumbers.Parse(number, "")
	if err != nil {
		panic(err)
	}
	fmt.Println("--------->", num.String())

	ok := phonenumbers.IsAlphaNumber(number)
	fmt.Println("---->", ok)

	ok = phonenumbers.IsValidNumber(num)
	fmt.Println("---->", ok)

	ok = phonenumbers.IsPossibleNumber(num)
	fmt.Println("---->", ok)

	// format it using national format
	formattedNum := phonenumbers.Format(num, phonenumbers.E164)

	fmt.Println("---->", num.GetCountryCode(), formattedNum)

	/*
		b := phonenumbers.GetRegionCodeForCountryCode(int(num.GetCountryCode()))
		fmt.Println("---->", b)
	*/
}
