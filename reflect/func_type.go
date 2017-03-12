package main

import (
	"fmt"
	"reflect"
)

type A int

type B struct {
	A
}

func (A) Av()  {}
func (*A) Ap() {}

func (B) Bv()  {}
func (*B) Bp() {}

func main() {
	var b B

	t := reflect.TypeOf(&b)
	s := []reflect.Type{t, t.Elem()}

	for _, x := range s {
		fmt.Println(x, ":", x.NumMethod())
		for i := 0; i < x.NumMethod(); i++ {
			fmt.Println("    ", x.Method(i))
		}
	}
}
