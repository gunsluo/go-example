package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var (
		wg  sync.WaitGroup
		ch  = make(chan int, 100)
		tch = make(chan struct{}, 10)
	)

	wg.Add(302)
	go func() {
		for i := 0; i < 300; i++ {
			reqeust(ch, i)
			time.Sleep(100 * time.Microsecond)
		}

		fmt.Println("send ok.")
		close(ch)
		wg.Done()
	}()

	go func() {
		task(&wg, ch, tch)
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("task over")
}

func reqeust(ch chan int, i int) {
	ch <- i
}

func task(wg *sync.WaitGroup, ch chan int, tch chan struct{}) {
	for i := range ch {
		tch <- struct{}{}
		go func(i int) {
			fmt.Println("->", i)
			time.Sleep(time.Second)
			<-tch
			wg.Done()
		}(i)
	}
}
