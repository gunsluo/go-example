package main

import (
	"fmt"

	lru "github.com/hashicorp/golang-lru"
)

func main() {
	l, err := lru.New(128)
	if err != nil {
		panic(err)
	}

	for i := 0; i < 256; i++ {
		l.Add(i, i*10)
	}
	if l.Len() != 128 {
		panic(fmt.Sprintf("bad len: %v", l.Len()))
	}

	if val, ok := l.Get(200); ok {
		fmt.Println("value: ", val)
	} else {
		panic("failed to get val")
	}

	if _, ok := l.Get(100); ok {
		panic("val not be remove")
	}
}
