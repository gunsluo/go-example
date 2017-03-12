package main

import (
	"fmt"
	"reflect"
)

func test() {
	x := 100

	tx, tp := reflect.TypeOf(x), reflect.TypeOf(&x)

	fmt.Println(tx, tp, tx == tp)
	fmt.Println(tx.Kind(), tp.Kind(), tp.Kind() == reflect.Ptr)
	fmt.Println(tx, tp.Elem())
}

func test2() {
	fmt.Println(reflect.TypeOf(map[string]int{}).Elem())
	fmt.Println(reflect.TypeOf([]int32{}).Elem())
}

func main() {
	test()
	test2()
}
