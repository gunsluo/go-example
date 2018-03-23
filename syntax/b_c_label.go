package main

import (
	"fmt"
)

func main() {
	breaklabel()
	continuelabel()
}

func breaklabel() {
	FirstNames := []string{"aaa", "bbb", "ccc"}
	LastNames := []string{"111", "222", "333"}

Loop:
	for _, firstName := range FirstNames {
		for _, lastName := range LastNames {
			fmt.Printf("Name: %s %s\n", firstName, lastName)

			if firstName == "bbb" && lastName == "111" {
				break Loop
			}
		}
	}
	println("Over.")
}

func continuelabel() {
	FirstNames := []string{"aaa", "bbb", "ccc"}
	LastNames := []string{"111", "222", "333"}

Loop:
	for _, firstName := range FirstNames {
		for _, lastName := range LastNames {
			fmt.Printf("Name: %s %s\n", firstName, lastName)

			if firstName == "bbb" && lastName == "111" {
				continue Loop
			}
		}
	}
	println("Over.")
}
