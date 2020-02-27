package main

import "fmt"

//go:generate msgp

type MessageType int

// MessageRouter is message router
type MessageRouter struct {
	Type         MessageType `msg:"type"`
	RoutingPaths []string    `msg:"routingPaths"`
	ClientIDs    []string    `msg:"clientId"`
	Namespace    string      `msg:"namespace"`
}

// Message is the message
type Message struct {
	Router MessageRouter `msg:"router"`
	Msg    []byte        `msg:"msg"`
}

func main() {
	msg := &Message{
		Router: MessageRouter{
			Type:         1,
			RoutingPaths: []string{"abc"},
		},
		Msg: []byte("hello"),
	}

	buffer, err := msg.MarshalMsg(nil)
	if err != nil {
		panic(err)
	}

	nmsg := &Message{}
	left, err := nmsg.UnmarshalMsg(buffer)
	if err != nil {
		panic(err)
	}

	fmt.Println("buffer->", string(left))
	fmt.Println("msg->", string(nmsg.Msg))
	fmt.Println("msg.Router->", nmsg.Router.Type, nmsg.Router.RoutingPaths)
}
