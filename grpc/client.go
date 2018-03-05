package main

import (
	"log"
	"os"
	"time"

	"github.com/gunsluo/go-example/grpc/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	//conn, err := grpc.Dial(address, grpc.WithInsecure())
	conn, err := grpc.Dial(address, grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(UnaryClientInterceptor),
		grpc.WithStreamInterceptor(StreamClientInterceptor))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}

func UnaryClientInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	log.Printf("before invoker. method: %+v, request:%+v", method, req)
	err := invoker(ctx, method, req, reply, cc, opts...)
	log.Printf("after invoker. reply: %+v", reply)
	return err
}

func StreamClientInterceptor(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	log.Printf("before invoker. method: %+v, StreamDesc:%+v", method, desc)
	clientStream, err := streamer(ctx, desc, cc, method, opts...)
	log.Printf("before invoker. method: %+v", method)
	return clientStream, err
}
