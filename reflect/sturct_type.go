package main

import (
	"fmt"
	"reflect"
)

type user struct {
	name string
	age  int
}

type manager struct {
	user
	title string
}

func main() {
	var m manager
	t := reflect.TypeOf(&m)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Println(f.Name, f.Type, f.Offset)

		if f.Anonymous {
			for x := 0; x < f.Type.NumField(); x++ {
				f1 := f.Type.Field(x)
				fmt.Println(f1.Name, f1.Type, f1.Offset)
			}
		}
	}
}
