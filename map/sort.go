package main

import (
	"fmt"
)

func main() {
	m := make(map[string]string)
	m["hello"] = "echo hello"
	m["world"] = "echo world"
	m["go"] = "echo go"
	m["is"] = "echo is"
	m["cool"] = "echo cool"

	values := make([]string, len(m))
	i := 0
	for k, v := range m {
		fmt.Printf("k=%v, v=%v\n", k, v)
		values[i] = v
	}

	fmt.Println("=>", values)
}
