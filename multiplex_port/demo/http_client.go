package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	resp, err := http.Post("http://127.0.0.1:50051/v1/example/echo",
		"application/json", strings.NewReader(`{"name": "world"}`))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("reply:", string(body))
}
