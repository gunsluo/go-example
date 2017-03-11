package main

import (
	"runtime"
	"time"
)

func test() {
	c := make(chan int)

	for i := 0; i < 10; i++ {
		go func() {
			<-c
		}()
	}
}

func main() {
	//GODEBUG="gctrace=1,schedtrace=1000,scheddetail=1" ./leak
	test()

	for {
		time.Sleep(time.Second)
		runtime.GC()
	}
}
