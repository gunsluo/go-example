package main

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/gunsluo/go-example/grpc/keepalive/pb"
	"google.golang.org/grpc"
)

func main() {
	gSrv := NewG()

	gSrv.Run()
}

type gServer struct {
	pb.UnimplementedGreeterServer

	port int
	gRPC *grpc.Server
}

func NewG() *gServer {
	gRPC := grpc.NewServer(
	// grpc.KeepaliveParams(keepalive.ServerParameters{
	// 	// MaxConnectionIdle:     24 * time.Hour,   // The current default value is infinity.
	// 	MaxConnectionAge:      2 * time.Minute,  // The current default value is infinity.
	// 	MaxConnectionAgeGrace: 30 * time.Second, // The current default value is infinity.
	// 	// Time:                  2 * time.Hour,    // The current default value is 2 hours.
	// 	// Timeout:               20 * time.Second, // The current default value is 20 seconds.
	// }),
	// grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
	// 	MinTime:             time.Minute, // The current default value is 5 minutes.
	// 	PermitWithoutStream: false,       // false by default.
	// }),
	)

	s := &gServer{
		port: 30000,
		gRPC: gRPC,
	}
	pb.RegisterGreeterServer(gRPC, s)

	return s
}

func (s *gServer) Run() {
	fmt.Printf("Starting gRPC server on listening %d.\n", s.port)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		panic(err)
	}

	s.gRPC.Serve(listener)
}

func (s *gServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	st := time.Now()
	time.Sleep(time.Minute)
	fmt.Println("say hello:", time.Now().Sub(st))
	return &pb.HelloReply{Message: "hello" + req.Name}, nil
}

func (s *gServer) SayHelloProgress(req *pb.HelloRequest, srv pb.Greeter_SayHelloProgressServer) error {
	st := time.Now()
	i := 0
	for i < 3 {
		time.Sleep(30 * time.Second)
		if err := srv.Send(&pb.SayHelloProgressReply{Progress: "processing"}); err != nil {
			fmt.Println("err:", time.Now().Sub(st), err)
			return err
		}

		i++
		fmt.Println("is progressing:", i)
	}

	// should be asynchronization
	if err := srv.Send(&pb.SayHelloProgressReply{Message: "hello" + req.Name, Progress: "done"}); err != nil {
		fmt.Println("err:", time.Now().Sub(st), err)
		return err
	}
	i++

	fmt.Println("say hello:", i, time.Now().Sub(st))
	return nil
}
