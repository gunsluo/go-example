package main

import (
	"github.com/nats-io/nats.go"
)

func main() {
	// Connect to a server
	nc, err := nats.Connect(nats.DefaultURL, nats.UserInfo("root", "password"))
	if err != nil {
		panic(err)
	}

	// Simple Publisher
	if err := nc.Publish("foo", []byte("Hello World")); err != nil {
		panic(err)
	}

	nc.Close()
}
