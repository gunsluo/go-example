package main

import (
	"fmt"

	"github.com/nyaruka/phonenumbers"
)

func main() {
	number := "+8618980501737"
	// 6502530000
	// parse our phone number
	num, err := phonenumbers.Parse(number, "")
	if err != nil {
		panic(err)
	}

	// format it using national format
	formattedNum := phonenumbers.Format(num, phonenumbers.NATIONAL)

	fmt.Println("---->", num.GetCountryCode(), formattedNum)

	b := phonenumbers.GetRegionCodeForCountryCode(int(num.GetCountryCode()))
	fmt.Println("---->", b)
}
