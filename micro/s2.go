package main

import (
	"fmt"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/broker"
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
		file.WithPath("config/s2.yaml"),
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
		micro.Name("s2"),
		micro.Address(cfg.Address),
	)

	// initialise flags
	service.Init()

	broker := service.Server().Options().Broker
	if err := broker.Connect(); err != nil {
		panic(err)
	}

	if _, err := broker.Subscribe("topic.s2", customSubscriber); err != nil {
		panic(err)
	}
	/*
		// register subscriber
		if err := micro.RegisterSubscriber("topic.s2", service.Server(), new(Sub)); err != nil {
			//if err := micro.RegisterSubscriber("topic.s2", service.Server(), subEv); err != nil {
			panic(err)
		}
	*/

	// start the service
	if err := service.Run(); err != nil {
		panic(err)
	}
}

/*
type Sub struct{}

func (s *Sub) Process(ctx context.Context, frame *proto.Event) error {
	md, _ := metadata.FromContext(ctx)
	fmt.Printf("[pubsub] Received frame %+v with metadata %+v\n", frame, md)
	// do something with event
	return nil
}

func subEv(ctx context.Context, frame *proto.Event) error {
	md, _ := metadata.FromContext(ctx)
	fmt.Printf("[pubsub] Received frame %+v with metadata %+v\n", frame, md)
	// do something with event
	return nil
}
*/

func customSubscriber(e broker.Event) error {
	if e == nil {
		return nil
	}

	msg := e.Message()
	fmt.Printf("[pubsub] Received broker event %+v %+v %+v %+v\n", e.Topic(), e.Error(), msg.Header, string(msg.Body))

	return nil
}
