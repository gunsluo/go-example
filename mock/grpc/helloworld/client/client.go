package main

import (
	"context"

	"github.com/gunsluo/go-example/grpc/helloworld/pb"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:19000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	greeterClient := pb.NewGreeterClient(conn)

	msg, err := sayHello(greeterClient, "luoji")
	if err != nil {
		logrus.WithError(err).Fatal("unable to sayhello")
	}

	logrus.Info("reply:", msg)
}

func sayHello(client pb.GreeterClient, name string) (string, error) {
	reply, err := client.SayHello(context.Background(),
		&pb.HelloRequest{Name: name})
	if err != nil {
		return "", err
	}

	return reply.Message, nil
}
