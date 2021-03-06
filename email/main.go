package main

import (
	"log"

	trumail "github.com/sdwolfe32/trumail/verifier"
)

func main() {
	v := trumail.NewVerifier("localhost", "no-reply@target-energysolutions.com")

	// Validate a single address
	log.Println(v.Verify("gunsluo@gmail.com"))
	//"ji.luo@target-energysolutions.com"
	log.Println(v.Verify("ji.luo@target-energysolutions.com"))
}
