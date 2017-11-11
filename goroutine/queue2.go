package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var (
		wg sync.WaitGroup
		ch = make(chan int, 1000)
	)

	wg.Add(102)
	go func() {
		for i := 0; i < 100; i++ {
			reqeust(ch, i)
			//time.Sleep(100 * time.Microsecond)
		}

		close(ch)
		fmt.Println("send ok.")
		wg.Done()
	}()

	go func() {
		task(&wg, ch, 10)
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("over")
}

func reqeust(ch chan int, i int) {
	ch <- i
}

func task(wg *sync.WaitGroup, ch chan int, cn int32) {
	var (
		c  int32
		c2 int32
		s  = make(chan struct{})
	)

	for i := range ch {
		//do something...
		go func(i int) {
			time.Sleep(time.Second)
			fmt.Println("->", i)
			wg.Done()

			atomic.AddInt32(&c2, 1)
			if c2 == cn {
				s <- struct{}{}
				atomic.StoreInt32(&c2, 0)
			}
		}(i)

		c++
		if c == cn {
			fmt.Println("->")
			<-s
			c = 0
		}
	}

	fmt.Println("task over")
}
