package main

import (
	"context"

	"github.com/gunsluo/go-example/multiplex_port/pb"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	greeterClient := pb.NewGreeterClient(conn)

	reply, err := greeterClient.SayHello(context.Background(),
		&pb.HelloRequest{Name: "luoji"})
	if err != nil {
		logrus.WithError(err).Fatal("unable to sayhello")
	}

	logrus.Info("reply:", reply.Message)
}
