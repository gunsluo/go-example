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

func process(r Request) {
	fmt.Println("process:", r)
}

func Serve(queue chan Request) {
	for req := range queue {
		<-sem
		/*
			go func() {
				// req is share resource
				process(req)
				sem <- 1
			}()
		*/
		// fix bug
		go func(req Request) {
			// person param
			process(req)
			sem <- 1
		}(req)
	}

	time.Sleep(2 * time.Second)
}

func main() {
	queue := make(chan Request)

	go func() {
		for i := 0; i < MaxOutstanding; i++ {
			queue <- Request(i)
		}
		close(queue)
	}()

	Serve(queue)
}
