package main

import (
	"fmt"

	"github.com/lithammer/shortuuid/v3"
)

func main() {
	u := shortuuid.New()   // Cekw67uyMpBGZLRP2HFVbe
	fmt.Println(len(u), u) // 6R7VqaQHbzC1xwA5UueGe6

	var total = 1000000
	var j int
	dm := make(map[string]interface{}, total)
	for i := 0; i < total; i++ {
		u := shortuuid.New() // Cekw67uyMpBGZLRP2HFVbe
		if _, ok := dm[u]; ok {
			fmt.Println("-->", u)
			j++
		} else {
			dm[u] = struct{}{}
		}
	}

	fmt.Println("->", total, j)
}
