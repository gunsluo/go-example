package main

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	acpb "tespkg.in/genproto/ac"
	configuratorpb "tespkg.in/genproto/configurator"
)

func main() {
	// testAC()
	testCfg()
}

func testAC() {
	conn, err := grpc.Dial("127.0.0.1:20000",
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			// Time:                time.Hour,        // The current default value is infinity.
			Time:    30 * time.Second, // The current default value is infinity.
			Timeout: 30 * time.Second, // The current default value is 20 seconds.
			// PermitWithoutStream: true, // false by default.
		}),
		// grpc.WithUnaryInterceptor(grpctrace.UnaryClientInterceptor()),
	)
	if err != nil {
		panic(err)
	}

	acClient := acpb.NewAccessControlClient(conn)
	ac(acClient)
}

func ac(client acpb.AccessControlClient) {
	ctx := context.Background()

	st := time.Now()
	resp, err := client.Ping(ctx, &acpb.PingRequest{Ping: "ping"})
	if err != nil {
		fmt.Println("err:", time.Now().Sub(st))
		panic(err)
	}
	fmt.Println("ping:", resp.Pong, time.Now().Sub(st))
}

func testCfg() {
	conn, err := grpc.Dial("127.0.0.1:21000",
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			// Time:                time.Hour,        // The current default value is infinity.
			Time:    10 * time.Second, // The current default value is infinity.
			Timeout: 10 * time.Second, // The current default value is 20 seconds.
			// PermitWithoutStream: true, // false by default.
		}),
		// grpc.WithUnaryInterceptor(grpctrace.UnaryClientInterceptor()),
	)
	if err != nil {
		panic(err)
	}

	cfgClient := configuratorpb.NewConfiguratorClient(conn)
	cfg(cfgClient)
}

func cfg(client configuratorpb.ConfiguratorClient) {
	ctx := context.Background()

	st := time.Now()
	resp, err := client.RegisterApp(ctx, &configuratorpb.RegisterAppRequest{})
	if err != nil {
		fmt.Println("err:", time.Now().Sub(st))
		panic(err)
	}
	fmt.Println("register app:", resp.Success, time.Now().Sub(st))
}
