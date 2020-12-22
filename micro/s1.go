package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/codec/bytes"
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
	// service.Server()

	// create publisher
	// p := micro.NewEvent("topic.s2", service.Client())
	pub2s2 := micro.NewPublisher("topic.s2", service.Client())

	go func() {
		time.Sleep(5 * time.Second)

		e := &Event{Message: "event"}
		buffer, err := json.Marshal(e)
		if err != nil {
			panic(err)
		}

		if err := pub2s2.Publish(context.Background(), &bytes.Frame{Data: buffer}); err != nil {
			fmt.Println("failed to publish event:", err)
			return
		}
		fmt.Println("publish event")
	}()

	// start the service
	if err := service.Run(); err != nil {
		panic(err)
	}
}

type Event struct {
	// unique id
	Id string `json:"id,omitempty"`
	// unix timestamp
	Timestamp int64 `json:"timestamp,omitempty"`
	// message
	Message string `json:"message,omitempty"`
}
