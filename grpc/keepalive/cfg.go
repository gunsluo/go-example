package main

import (
	"context"
	"fmt"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	acpb "tespkg.in/genproto/ac"
	configuratorpb "tespkg.in/genproto/configurator"
)

func main() {
	cfgSrv := NewCfg()

	cfgSrv.Run()
}

type cfgServer struct {
	configuratorpb.UnimplementedConfiguratorServer

	port int
	gRPC *grpc.Server

	acClient acpb.AccessControlClient
}

func NewCfg() *cfgServer {
	conn, err := grpc.Dial("127.0.0.1:20000",
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			// Time:                time.Hour,        // The current default value is infinity.
			Time:    20 * time.Second, // The current default value is infinity.
			Timeout: 10 * time.Second, // The current default value is 20 seconds.
			// PermitWithoutStream: true, // false by default.
		}),
	)
	if err != nil {
		panic(err)
	}
	acClient := acpb.NewAccessControlClient(conn)

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

	s := &cfgServer{
		port:     21000,
		gRPC:     gRPC,
		acClient: acClient,
	}
	configuratorpb.RegisterConfiguratorServer(gRPC, s)

	return s
}

func (s *cfgServer) Run() {
	fmt.Printf("Starting gRPC server on listening %d.\n", s.port)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		panic(err)
	}

	s.gRPC.Serve(listener)
}

func (s *cfgServer) RegisterApp(ctx context.Context, req *configuratorpb.RegisterAppRequest) (*configuratorpb.RegisterAppReply, error) {
	st := time.Now()
	_, err := s.acClient.Ping(ctx, &acpb.PingRequest{Ping: "ping"})
	if err != nil {
		fmt.Println("--->", time.Now().Sub(st), err)
		return nil, err
	}
	fmt.Println("RegisterApp:", time.Now().Sub(st))

	emptyRsp := &configuratorpb.RegisterAppReply{}
	emptyRsp.Success = true
	return emptyRsp, nil
}
