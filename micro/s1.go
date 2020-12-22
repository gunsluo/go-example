package main

import (
	"fmt"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/encoder/yaml"
	"github.com/micro/go-micro/v2/config/source"
	"github.com/micro/go-micro/v2/config/source/file"
)

func main() {
	var cfg struct {
		Address string `json:"address"`
	}

	enc := yaml.NewEncoder()
	config.Load(file.NewSource(
		file.WithPath("config/s1.yaml"),
		source.WithEncoder(enc),
	))

	// fmt.Printf("---->%#v", config.Map())
	// var address string
	if err := config.Get().Scan(&cfg); err != nil {
		panic(err)
	}
	fmt.Println(cfg.Address)

	w, err := config.Watch()
	if err != nil {
		panic(err)
	}
	go func() {
		for {
			v, err := w.Next()
			if err != nil {
				panic(err)
			}

			if err := v.Scan(&cfg); err != nil {
				panic(err)
			}
			fmt.Println("watch config", cfg.Address)
		}
	}()

	// create a new service
	service := micro.NewService(
		micro.Name("s1"),
		micro.Address(cfg.Address),
	)

	// initialise flags
	service.Init()

	service.Server()

	// start the service
	service.Run()
}
