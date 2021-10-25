package main

import (
	"embed"
	"fmt"
)

//go:embed hello.txt
var s string

//go:embed hello.txt
var b []byte

//go:embed hello.txt
var f embed.FS

func main() {
	fmt.Println(s)
	fmt.Println(string(b))

	data, _ := f.ReadFile("hello.txt")
	fmt.Println(string(data))
}
