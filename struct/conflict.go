package main

import "fmt"

type R struct {
	x int
}

func (r *R) Read() {
	fmt.Printf("read %T\n", r)
}

type W struct {
}

func (w *W) Write() {
	fmt.Printf("write %T\n", w)
}

type RW struct {
	x int
	R
	W
}

func (rw *RW) Read() {
	fmt.Printf("rw read %T\n", rw)
}

func (rw *RW) ReadWrite() {
	fmt.Printf("readwrite %T\n", rw)
}

func main() {
	rw := new(RW)
	rw.Read()
	rw.Write()
	rw.ReadWrite()
}
