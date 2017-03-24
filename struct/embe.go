package main

import "fmt"

type R struct {
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
	R
	W
}

func (rw *RW) ReadWrite() {
	fmt.Printf("readwrite %T\n", rw)
}

type RW2 struct {
	*R
	*W
}

func (rw *RW2) ReadWrite() {
	fmt.Printf("readwrite %T\n", rw)
}

func main() {
	rw := new(RW)
	rw.Read()
	rw.Write()
	rw.ReadWrite()

	rw2 := new(RW2)
	rw2.Read()
	rw2.Write()
	rw2.ReadWrite()
}
