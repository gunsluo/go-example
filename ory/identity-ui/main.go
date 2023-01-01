package main

import "github.com/gunsluo/go-example/ory/identity-ui/srv"

func main() {
	s, err := srv.New()
	if err != nil {
		panic(err)
	}

	s.Run()
}
