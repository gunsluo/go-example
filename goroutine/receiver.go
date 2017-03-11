package main

import "sync"

type receiver struct {
	sync.WaitGroup
	data   chan int
	Action func(m int)
}

func (rev *receiver) Send(m int) {
	rev.data <- m
}

func (rev *receiver) Close() {
	close(rev.data)
}

func newReceiver() *receiver {
	rev := &receiver{
		data: make(chan int),
	}

	rev.Add(1)
	go func() {
		defer rev.Done()
		for x := range rev.data {
			rev.Action(x)
		}

	}()

	return rev
}

func main() {
	rev := newReceiver()
	rev.Action = func(m int) {
		println(m)
	}
	rev.Send(1)
	rev.Send(2)

	rev.Close()
	rev.Wait()
}
