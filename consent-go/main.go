package main

import "github.com/gunsluo/go-example/consent-go/srv"

func main() {
	s, err := srv.New()
	if err != nil {
		panic(err)
	}

	s.Run()
}
