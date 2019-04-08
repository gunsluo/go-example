package main

import (
	"fmt"
	"time"

	"github.com/jasonlvhit/gocron"
)

func task() {
	fmt.Println("I am runnning task.", time.Now())
}

func taskWithParams(a int, b string) {
	fmt.Println(a, b)
}

func main() {
	// also , you can create a your new scheduler,
	// to run two scheduler concurrently
	s := gocron.NewScheduler()
	s.Every(1).Day().At("00:16").Do(task)
	<-s.Start()
}
