package main

import (
	"context"
	"fmt"
	"net"
	"os"

	"github.com/gunsluo/go-example/grpc/ssl/pb"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

var (
	argAddress string
	argCrtFile string
	argKeyFile string
)

var verbose bool
var rootCmd *cobra.Command

func init() {
	rootCmd = &cobra.Command{
		Use:   "grpc",
		Short: "demo service",
		Long:  "Top level command for demo service, it provides GRPC service",
		Run:   run,
	}

	rootCmd.Flags().StringVarP(&argAddress, "address", "a", ":3264", "address to listen on")
	rootCmd.Flags().StringVar(&argCrtFile, "cert-file", "", "certificate file for gRPC TLS authentication")
	rootCmd.Flags().StringVar(&argKeyFile, "key-file", "", "key file for gRPC TLS authentication")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
}

type Service struct {
}

func (s *Service) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, grpc.Errorf(codes.Unauthenticated, "no metadata")
	}

	token := md.Get("token")
	if len(token) == 0 {
		return nil, grpc.Errorf(codes.Unauthenticated, "no token")
	}

	fmt.Println("requst:", token[0], req.Name)
	return &pb.HelloReply{
		Message: "hello, " + req.Name,
	}, nil
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run(cmd *cobra.Command, _ []string) {
	listener, err := net.Listen("tcp", argAddress)
	if err != nil {
		panic(err)
	}

	// Create the TLS credentials
	var opts []grpc.ServerOption
	if argCrtFile != "" && argKeyFile != "" {
		fmt.Println("enable credentials in the grpc")
		creds, err := credentials.NewServerTLSFromFile(argCrtFile, argKeyFile)
		if err != nil {
			panic(err)
		}

		opts = append(opts, grpc.Creds(creds))
	}

	server := grpc.NewServer(opts...)
	pb.RegisterGreeterServer(server, &Service{})

	logrus.WithField("addr", argAddress).Println("Starting server")
	server.Serve(listener)
}
