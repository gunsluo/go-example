package main

import (
	"fmt"
	"unsafe"
)

func main() {
	s := make([]byte, 200)
	ptr := unsafe.Pointer(&s[0])
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", ptr)

	s1 := ((*[1 << 10]byte)(ptr))[:200]
	fmt.Printf("%T, %d, %d\n", s1, len(s1), cap(s1))

	var sl = struct {
		addr uintptr
		len  int
		cap  int
	}{uintptr(ptr), 200, 512}
	s2 := *(*[]byte)(unsafe.Pointer(&sl))
	fmt.Printf("%T, %d, %d\n", s2, len(s2), cap(s2))
}
