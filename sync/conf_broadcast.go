package main

import (
	"fmt"
	"sync"
	"time"
)

func work(cond *sync.Cond, i int) {
	fmt.Println("work", i)
	cond.L.Lock()
	cond.Wait()
	cond.L.Unlock()
}

func main() {
	var cond = sync.NewCond(new(sync.Mutex))
	for i := 0; i < 10; i++ {
		go work(cond, i)
	}
	// 下面的 sleep 很重要
	time.Sleep(time.Duration(2) * time.Second)
	fmt.Println("Wake up")
	cond.Broadcast()
	//time.Sleep(time.Duration(20) * time.Second)
	fmt.Println("Done")
}
