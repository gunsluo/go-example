package main

import (
	"fmt"
	"reflect"
)

func main() {
	type user struct {
		Name string
		Age  int
	}

	u := user{
		"q.yuhen",
		60,
	}

	v := reflect.ValueOf(&u)

	if !v.CanInterface() {
		println("CanInterface: fail.")
		return
	}

	p, ok := v.Interface().(*user)
	if !ok {
		println("Interface: fail.")
		return
	}

	p.Age++
	fmt.Printf("%+v\n", p)
}
