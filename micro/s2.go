package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/config"
)

func main() {
	// create a new service
	service := micro.NewService(
		micro.Name("helloworld"),
		micro.Address(":10000"),
	)

	// initialise flags
	service.Init()

	service.Server()

	// start the service
	service.Run()

	w, err := config.Watch("hosts", "database")

	v, err := w.Next()
}
