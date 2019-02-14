package main

import (
	"context"
	"fmt"
	"net"

	"github.com/gunsluo/go-example/grpc/helloworld/pb"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const (
	sAddress = "0.0.0.0:9000"
)

type Service struct {
	requestNum int
}

func (s *Service) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	fmt.Println("reveive:", req.Name)
	s.requestNum++
	return &pb.HelloReply{
		Message: fmt.Sprintf("+%d hello, %s", s.requestNum, req.Name),
	}, nil
}

func main() {
	server := grpc.NewServer()
	pb.RegisterGreeterServer(server, &Service{})

	listener, err := net.Listen("tcp", sAddress)
	if err != nil {
		panic(err)
	}

	logrus.WithField("addr", sAddress).Println("Starting server")
	server.Serve(listener)
}
