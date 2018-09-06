package main

import (
	"context"
	"fmt"

	"github.com/gunsluo/go-example/grpc/ssl/pb"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	sAddress = "0.0.0.0:3264"
	crtFile  = "../cert/server.crt"
)

func main() {
	var opts []grpc.DialOption
	if crtFile != "" {
		fmt.Println("enable credentials in the grpc")
		// target is common name(host name) in the cert file
		creds, err := credentials.NewClientTLSFromFile(crtFile, "target")
		if err != nil {
			panic(err)
		}

		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}

	conn, err := grpc.Dial("127.0.0.1:3264", opts...)
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
	conn.Close()
}
