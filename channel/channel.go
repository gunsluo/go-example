package main

import (
	"fmt"
)

type Request struct {
	args []int
	f    func([]int) int
	rc   chan int
}

func sum(a []int) (s int) {
	for _, v := range a {
		s += v
	}

	return
}

//rpc client
func client(conn chan *Request) {

	req := &Request{[]int{3, 4, 5}, sum, make(chan int)}
	conn <- req

	fmt.Println("response:", <-req.rc)
}

//rpc serve
func serve(queue chan *Request) {
	for req := range queue {
		req.rc <- req.f(req.args)
	}
}

func main() {
	conn := make(chan *Request)
	go serve(conn)

	client(conn)
}
