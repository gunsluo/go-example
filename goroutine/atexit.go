package main

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var exits = struct {
	sync.RWMutex
	funcs   []func()
	signals chan os.Signal
}{}

func atexit(fn func()) {
	exits.Lock()
	defer exits.Unlock()
	exits.funcs = append(exits.funcs, fn)
}

func waitexit() {
	if exits.signals == nil {
		exits.signals = make(chan os.Signal)
		signal.Notify(exits.signals, syscall.SIGINT, syscall.SIGTERM)
	}

	exits.RLock()
	for _, fn := range exits.funcs {
		defer fn()
	}
	exits.RUnlock()

	<-exits.signals
}

func main() {
	atexit(func() {
		println("exit one...")
	})

	atexit(func() {
		println("exit two...")
	})

	waitexit()
}
