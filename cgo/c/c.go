package main

import (
	"fmt"

	"github.com/gunsluo/go-example/cgo/c/add"
)

func main() {
	fmt.Println(add.Add(2, 15))
}
