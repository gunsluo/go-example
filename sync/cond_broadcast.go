package main

import (
	"fmt"
	"sync"
	"time"
)

func work(cond *sync.Cond, i int) {
	cond.L.Lock()
	cond.Wait()
	cond.L.Unlock()
	fmt.Println("work", i)
}

func main() {
	var wg sync.WaitGroup
	var cond = sync.NewCond(new(sync.Mutex))
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			cond.L.Lock()
			cond.Wait()
			cond.L.Unlock()
			fmt.Println("work", i)
			wg.Done()
		}(i)
	}
	// 下面的 sleep 很重要
	time.Sleep(time.Second)
	fmt.Println("Wake up")
	cond.Broadcast()
	//cond.Signal()
	wg.Wait()
	fmt.Println("Done")
}
