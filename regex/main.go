package main

import (
	"fmt"
	"regexp"

	"github.com/dlclark/regexp2"
)

func main() {
	{
		re, err := regexp.Compile(`(?:.*?[A-Z])(?:.*?[0-9])`)
		//re, err := regexp.Compile(`(?=.*?[A-Z])`)
		//re, err := regexp.Compile(`([A-Z,a-z]+`)
		//re, err := regexp.Compile(`(?!^[0-9]*$)(?!^[a-zA-Z]*$)^([a-zA-Z0-9]{6,50})$`)
		if err != nil {
			panic(err)
		}
		fmt.Println(re.MatchString("Gopher123"))
		fmt.Println(re.MatchString("Gophergopher123"))
		fmt.Println(re.MatchString("Gophergophergopher"))
	}
	{
		//re := regexp2.MustCompile(`(?=.*?[A-Z])(?=.*?[0-9])`, 0)
		re := regexp2.MustCompile(`Windows (?:95|98|NT|2000) abc`, 0)
		//re := regexp2.MustCompile(`Windows (?=95|98|NT|2000)[0-9]{2,4} abc`, 0)
		isMatch, err := re.MatchString(`Windows`)
		fmt.Println(isMatch, err)
		isMatch, err = re.MatchString(`Windows 95`)
		fmt.Println(isMatch, err)
		isMatch, err = re.MatchString(`Windows 3.1`)
		fmt.Println(isMatch, err)
		isMatch, err = re.MatchString(`Windows 95 abc`)
		fmt.Println(isMatch, err)
		isMatch, err = re.MatchString(`Windows 98 abc`)
		fmt.Println(isMatch, err)
	}
}
