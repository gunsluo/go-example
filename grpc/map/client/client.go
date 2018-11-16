package main

import (
	"context"

	"github.com/gunsluo/go-example/grpc/map/pb"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:19000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	demoClient := pb.NewDemoClient(conn)

	reply, err := demoClient.Hello(context.Background(),
		&pb.HelloRequest{
			Conditions: map[string][]byte{
				"options": []byte("Yes"),
			},
		})
	if err != nil {
		logrus.WithError(err).Fatal("unable to sayhello")
	}

	logrus.Info("reply:", reply.Message)
}
