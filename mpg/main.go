package main

import (
	"fmt"
	"sync"
	"time"
)

type G interface {
	Run()
}

type Sched struct {
	allg []G
	lock sync.Mutex
}

var sched Sched

func M() {
	for {
		sched.lock.Lock()
		if len(sched.allg) > 0 {
			g := sched.allg[0]
			sched.allg = sched.allg[1:]
			sched.lock.Unlock()
			g.Run()
		} else {
			sched.lock.Unlock()
		}
	}
}

const (
	GOMAXPROCS = 2
)

func main() {
	for i := 0; i < GOMAXPROCS; i++ {
		go M()
	}

	for {
		for i := 0; i < GOMAXPROCS; i++ {
			g := new(Gorutine)
			g.x = time.Now().Unix()
			sched.lock.Lock()
			sched.allg = append(sched.allg, g)
			sched.lock.Unlock()
		}
		time.Sleep(1 * time.Second)
	}
}

func entersyscall() {
	go M()
}

func exitsyscall() {
	/*
		if len(allm) >= GOMAXPROCS {
			sched.lock.Lock()
			sched.allg = append(sched.allg, g)
			sched.lock.Unlock()
			time.Sleep()
		}
	*/
}

type Gorutine struct {
	x int64
}

func (g *Gorutine) Run() {
	fmt.Println("G:", g.x)
}
