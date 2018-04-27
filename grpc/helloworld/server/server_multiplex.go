package main

import (
	"context"
	"fmt"
	"net"
	"strings"

	"github.com/gunsluo/go-example/grpc/helloworld/pb"
	"github.com/sirupsen/logrus"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
)

const (
	sAddress = "0.0.0.0:19000"
)

type Service struct {
}

func (s *Service) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	fmt.Println("==>", req.Name)
	return &pb.HelloReply{
		Message: "hello, " + req.Name,
	}, nil
}

func main() {
	server := grpc.NewServer()
	pb.RegisterGreeterServer(server, &Service{})

	listener, err := net.Listen("tcp", sAddress)
	if err != nil {
		logrus.WithError(err).Fatalf("Can't connect serve")
	}

	m := cmux.New(listener)
	//grpcl := m.Match(cmux.HTTP2HeaderFieldPrefix("content-type", "application/grpc"))
	grpcl := m.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))

	logrus.WithField("addr", sAddress).Println("Starting server")
	go server.Serve(grpcl)

	if err := m.Serve(); !strings.Contains(err.Error(), "use of closed network connection") {
		logrus.WithError(err).Fatalf("multiplex port serve")
	}
}
