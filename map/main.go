package main

import (
	"fmt"
	"sync"
)

func main() {
	m := &sync.Map{}
	v, ok := m.Load("key")
	fmt.Println(v, ok)
	v, ok = m.LoadOrStore("key", "val")
	fmt.Println(v, ok)
	v, ok = m.LoadOrStore("key", "val1")
	fmt.Println(v, ok)
	m.Store("key", "val1")
	v, ok = m.LoadOrStore("key", "val")
	fmt.Println(v, ok)

	m.Store("key1", "val")
	m.Range(func(k, v interface{}) bool {
		fmt.Println(k, v)
		return true
	})
}
