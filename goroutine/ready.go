package main

import (
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ready := make(chan int)

	for i := 0; i < 3; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			println("ready:", i)
			<-ready
			println("running:", i)
		}(i)
	}

	time.Sleep(1 * time.Second)
	println("ready Go?")
	close(ready)
	wg.Wait()
}
