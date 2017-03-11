package main

import (
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	one := make(chan int)
	two := make(chan int)

	go func() {
		defer wg.Done()
		for {
			select {
			case x, ok := <-one:
				if !ok {
					one = nil
					break
				}
				println("one:", x)
			case x, ok := <-two:
				if !ok {
					two = nil
					break
				}
				println("two:", x)
			}

			if one == nil && two == nil {
				return
			}
		}
	}()

	go func() {
		defer wg.Done()
		defer close(one)
		one <- 1
	}()

	go func() {
		defer wg.Done()
		defer close(two)
		two <- 2
	}()

	wg.Wait()
}
