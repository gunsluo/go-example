package main

import (
	"context"

	"github.com/gunsluo/go-example/grpc/helloworld/pb"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:19000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	greeterClient := pb.NewGreeterClient(conn)

	md := metadata.Pairs("authorization", "bearer abcd")
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	reply, err := greeterClient.SayHello(ctx,
		&pb.HelloRequest{Name: "luoji"})
	if err != nil {
		logrus.WithError(err).Fatal("unable to sayhello")
	}

	logrus.Info("reply:", reply.Message)
}
