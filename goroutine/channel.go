package main

import (
	"fmt"
	"time"
)

const (
	MaxOutstanding = 5
)

var (
	sem = make(chan int, MaxOutstanding)
)

type Request int

func init() {
	for i := 0; i < MaxOutstanding; i++ {
		sem <- 1
	}
}

func handle(r Request) {
	<-sem
	process(r)
	sem <- 1
}

func process(r Request) {
	fmt.Println("process:", r)
}

func Serve(queue chan Request) {
	for {
		req, ok := <-queue
		if ok {
			go handle(req)
		} else {
			break
		}
	}

	time.Sleep(2 * time.Second)
}

func main() {
	queue := make(chan Request)

	go func() {
		for i := 0; i < MaxOutstanding; i++ {
			time.Sleep(1 * time.Second)
			queue <- Request(i)
		}
		close(queue)
	}()

	Serve(queue)
}
