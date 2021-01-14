package main

import (
	"fmt"
	"sync"
	"time"
)

var i int

func sendTask(tasksCh chan int, wg *sync.WaitGroup) {
	for {
		lock.RLock()
		if over {
			lock.RUnlock()
			break
		}
		lock.RUnlock()

		i++
		wg.Add(1)
		fmt.Println("send", i)
		tasksCh <- i
		time.Sleep(time.Second)
	}
}

func start(tasksCh chan int, exit <-chan struct{}, wg *sync.WaitGroup) {
	for {
		select {
		case i := <-tasksCh:
			go func(i int) {
				time.Sleep(time.Second * 2)
				fmt.Println("processing task", i)
				wg.Done()
			}(i)
		case <-exit:
			wg.Wait()
			return
		}
	}
}

var over bool
var lock sync.RWMutex

func shutdown(exit chan<- struct{}, d time.Duration) {
	time.Sleep(d)
	lock.Lock()
	over = true
	lock.Unlock()
	exit <- struct{}{}
}

func main() {
	var wg sync.WaitGroup
	tasksCh := make(chan int)
	exit := make(chan struct{})

	go sendTask(tasksCh, &wg)
	go shutdown(exit, time.Second*10)

	start(tasksCh, exit, &wg)

	fmt.Println("over")
}
