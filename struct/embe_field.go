package main

import "fmt"

type R struct {
	x int
}

func (r *R) Read() {
	fmt.Printf("read %T x %d\n", r, r.x)
}

type W struct {
	x int
	y int
}

func (w *W) Write() {
	fmt.Printf("write %T x %d y %d\n", w, w.x, w.y)
}

type RW struct {
	y int
	z int
	R
	W
}

func (rw *RW) ReadWrite() {
	//fmt.Printf("rw readWrite %T x %d y %d z %d\n", rw, rw.x, rw.y, rw.z)
	fmt.Printf("rw readWrite %T y %d z %d\n", rw, rw.y, rw.z)
}

func main() {
	rw := new(RW)
	//rw.x = 1
	rw.y = 2
	rw.z = 3
	rw.Read()
	rw.Write()
	rw.ReadWrite()
}
