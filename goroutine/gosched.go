package main

import "runtime"

func main() {
	runtime.GOMAXPROCS(1)
	exit := make(chan int)

	go func() {
		defer close(exit)

		go func() {
			println("b")
		}()

		go func() {
			println("c")
		}()

		for i := 0; i < 4; i++ {
			println("a:", i)
			if i == 1 || i == 2 {
				runtime.Gosched()
			}
		}
	}()

	<-exit
}
