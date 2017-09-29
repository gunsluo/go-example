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
		gn = 2
	)

	wg.Add(3)
	go func() {
		for i := 0; i < 50; i++ {
			reqeust(ch, i)
			time.Sleep(100 * time.Microsecond)
		}

		fmt.Println("send ok.")
		close(ch)
		wg.Done()
	}()

	group(&wg, ch, gn)

	wg.Wait()
	fmt.Println("task over")
}

func reqeust(ch chan int, i int) {
	ch <- i
}

func group(wg *sync.WaitGroup, ch chan int, gn int) {
	queues := make([]chan int, gn)
	for i := 0; i < gn; i++ {
		queues[i] = make(chan int, 1)
		go func(i int) {
			for d := range queues[i] {
				fmt.Println("->", d)
				time.Sleep(time.Second)
			}
			wg.Done()
		}(i)
	}

	for i := range ch {
		queues[i%2] <- i
	}

	for i := 0; i < gn; i++ {
		close(queues[i])
	}
}
