package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var (
		wg sync.WaitGroup
		ch = make(chan int, 10)
		tn = 2
	)

	wg.Add(52)
	go func() {
		for i := 0; i < 50; i++ {
			reqeust(ch, i)
			time.Sleep(100 * time.Microsecond)
		}

		fmt.Println("send ok.")
		close(ch)
		wg.Done()
	}()

	go func() {
		task(&wg, ch, tn)
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("task over")
}

func reqeust(ch chan int, i int) {
	ch <- i
}

func task(wg *sync.WaitGroup, ch chan int, tn int) {
	queues := make([]chan struct{}, tn)
	for i := 0; i < tn; i++ {
		queues[i] = make(chan struct{}, 1)
	}

	for i := range ch {
		qch := queues[i%2]
		qch <- struct{}{}
		go func(i int) {
			fmt.Println("->", i)
			time.Sleep(time.Second)
			<-qch
			wg.Done()
		}(i)
	}
}
