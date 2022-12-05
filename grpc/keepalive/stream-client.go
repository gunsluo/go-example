package main

import (
	"context"
	"fmt"
	"time"

	"github.com/gunsluo/go-example/grpc/keepalive/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:30000",
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

	gClient := pb.NewGreeterClient(conn)
	stream(gClient)
}

func g(client pb.GreeterClient) {
	ctx := context.Background()
	// client.PubHello()

	st := time.Now()
	resp, err := client.SayHello(ctx, &pb.HelloRequest{Name: "luoji"})
	if err != nil {
		fmt.Println("err:", time.Now().Sub(st))
		panic(err)
	}
	fmt.Println("say hello:", resp.Message, time.Now().Sub(st))
}

func stream(client pb.GreeterClient) {
	ctx := context.Background()

	st := time.Now()
	hStream, err := client.SayHelloProgress(ctx, &pb.HelloRequest{Name: "luoji"})
	if err != nil {
		fmt.Println("err:", time.Now().Sub(st))
		panic(err)
	}

	var resp *pb.SayHelloProgressReply
	var i int
	for {
		res, err := hStream.Recv()
		if err != nil {
			fmt.Println("recv err:", time.Now().Sub(st))
			panic(err)
		}

		i++
		fmt.Println("recv message:", i)
		if res.Progress == "done" {
			resp = res
			break
		}
	}

	fmt.Println("say hello:", i, resp.Message, time.Now().Sub(st))
}
