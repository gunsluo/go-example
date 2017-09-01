package main

import (
	"fmt"
	"time"
)

type Request struct {
	Data string
}

type Response struct {
	Request
	MetadataInfo string
}

type Signal struct {
	requestCh    chan Request
	responseCh   chan Response
	workerTokens chan bool
}

func NewSignal() *Signal {
	return &Signal{
		requestCh:  make(chan Request),
		responseCh: make(chan Response),
	}
}

func (s *Signal) SendSignal(data string) {
	s.requestCh <- Request{Data: data}
}

func (s *Signal) ReceivingSignal() <-chan Response {
	return s.responseCh
}

func (s *Signal) Run() {
	for r := range s.requestCh {
		go func(r Request) {
			s.responseCh <- Response{
				Request:      r,
				MetadataInfo: "this is test",
			}
		}(r)
	}
}

func main() {
	s := NewSignal()
	go s.Run()

	go func() {
		for resp := range s.ReceivingSignal() {
			fmt.Println("-->", resp)
		}
	}()

	go func() {
		timer := time.NewTimer(1 * time.Second)
		for {
			<-timer.C
			s.SendSignal("hello")
			timer.Reset(3 * time.Second)
		}
	}()

	select {}
}
