package main

import (
	"fmt"
	"sync"
)

func work(c *sync.Cond) {
	fmt.Println("Notify main")
	c.Signal()
}

func main() {
	var (
		locker = new(sync.Mutex)
		cond   = sync.NewCond(locker)
	)
	cond.L.Lock()
	go work(cond)
	cond.Wait()
	fmt.Println("Done")
}
