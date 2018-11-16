package main

import (
	"context"
	"fmt"
	"net"

	"github.com/gunsluo/go-example/grpc/map/pb"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const (
	sAddress = "0.0.0.0:19000"
)

type Service struct {
}

func (s *Service) Hello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	for key, val := range req.Conditions {
		fmt.Println("==>", key, string(val))
	}

	return &pb.HelloReply{
		Message: "Ok",
	}, nil
}

func main() {
	server := grpc.NewServer()
	pb.RegisterDemoServer(server, &Service{})

	listener, err := net.Listen("tcp", sAddress)
	if err != nil {
		panic(err)
	}

	logrus.WithField("addr", sAddress).Println("Starting server")
	server.Serve(listener)
}
