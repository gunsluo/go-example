package main

import (
	"fmt"
	"time"
)

var ticker *time.Ticker

func main() {
	go tick()

	time.Sleep(4 * time.Second)
	ticker.Stop()
}

func tick() {
	ticker = time.NewTicker(1 * time.Second)
	for {
		select {
		case <-ticker.C:
			do()
		}
	}
}

func do() {
	fmt.Println(time.Now(), ":do something, begin.")
	time.Sleep(2 * time.Second)
	fmt.Println(time.Now(), ":do something, spend 2s.")
}
