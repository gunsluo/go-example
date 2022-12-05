package main

import (
	"context"
	"fmt"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	acpb "tespkg.in/genproto/ac"
)

func main() {
	acSrv := NewAC()

	acSrv.Run()
}

type acServer struct {
	acpb.UnimplementedAccessControlServer

	port int
	gRPC *grpc.Server
}

func NewAC() *acServer {
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

	s := &acServer{
		port: 20000,
		gRPC: gRPC,
	}
	acpb.RegisterAccessControlServer(gRPC, s)

	return s
}

func (s *acServer) Run() {
	fmt.Printf("Starting gRPC server on listening %d.\n", s.port)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		panic(err)
	}

	s.gRPC.Serve(listener)
}

func (s *acServer) Ping(ctx context.Context, req *acpb.PingRequest) (*acpb.PongReply, error) {
	st := time.Now()
	if req == nil || req.Ping != "ping" {
		return &acpb.PongReply{
				Pong: "Invalid Parameter",
			},
			status.Error(codes.InvalidArgument, "expecting ping to be 'ping'")
	}

	time.Sleep(3 * time.Minute)
	fmt.Println("pong:", time.Now().Sub(st))
	return &acpb.PongReply{Pong: "pong"}, nil
}
