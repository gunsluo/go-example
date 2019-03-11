package main

import (
	"context"
	"fmt"
	"net"

	"github.com/gunsluo/go-example/grpc/helloworld/pb"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
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

func grpcRecovery(err interface{}) error {
	fmt.Println("Panic:", err)
	return nil
}

func grpcAuth(ctx context.Context) (context.Context, error) {
	token, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, err
	}
	// TODO: check token
	_ = token

	return ctx, nil
}

func main() {
	opts := []grpc.ServerOption{}
	recovery := grpc_recovery.WithRecoveryHandler(grpcRecovery)

	opts = append(opts, grpc_middleware.WithUnaryServerChain(
		grpc_recovery.UnaryServerInterceptor(recovery),
		grpc_auth.UnaryServerInterceptor(grpcAuth),
	))
	opts = append(opts, grpc_middleware.WithStreamServerChain(
		grpc_recovery.StreamServerInterceptor(recovery),
		grpc_auth.StreamServerInterceptor(grpcAuth),
	))

	server := grpc.NewServer(opts...)
	pb.RegisterGreeterServer(server, &Service{})

	listener, err := net.Listen("tcp", sAddress)
	if err != nil {
		panic(err)
	}

	logrus.WithField("addr", sAddress).Println("Starting server")
	server.Serve(listener)
}
