package main

import (
	"fmt"
	"runtime"

	"github.com/nats-io/nats.go"
)

func main() {
	// Connect to a server
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		panic(err)
	}

	// Simple Async Subscriber
	nc.Subscribe("foo", func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})

	nc.Flush()
	// Simple Sync Subscriber
	//sub, err := nc.SubscribeSync("foo")
	//m, err := sub.NextMsg(timeout)

	//nc.Drain()

	runtime.Goexit()
	//nc.Close()
}
