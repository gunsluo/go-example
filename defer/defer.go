package main

import "fmt"

func main() {
	fmt.Println("f():", f())
	fmt.Println("f1():", f1())
	fmt.Println("f2():", f2())
}

func f() (res int) {
	defer func() {
		res++
	}()

	return 0
}

func f1() (res int) {
	t := 5
	defer func() {
		t += 5
	}()

	return t
}

func f2() (res int) {
	defer func(res int) {
		res += 5
	}(res)

	return 1
}
