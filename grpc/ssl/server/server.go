package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
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
	argCAFile  string
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
	rootCmd.Flags().StringVar(&argCAFile, "ca-file", "", "ca file for gRPC client")
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

		if argCAFile == "" {
			creds, err := credentials.NewServerTLSFromFile(argCrtFile, argKeyFile)
			if err != nil {
				panic(err)
			}

			opts = append(opts, grpc.Creds(creds))
		} else {
			// Parse certificates from certificate file and key file for server.
			cert, err := tls.LoadX509KeyPair(argCrtFile, argKeyFile)
			if err != nil {
				panic(err)
				//return fmt.Errorf("invalid config: error parsing gRPC certificate file: %v", err)
			}

			// Parse certificates from client CA file to a new CertPool.
			cPool := x509.NewCertPool()
			clientCert, err := ioutil.ReadFile(argCAFile)
			if err != nil {
				panic(err)
				//return fmt.Errorf("invalid config: reading from client CA file: %v", err)
			}
			if cPool.AppendCertsFromPEM(clientCert) != true {
				panic(err)
				//return errors.New("invalid config: failed to parse client CA")
			}

			tlsConfig := tls.Config{
				Certificates: []tls.Certificate{cert},
				ClientAuth:   tls.RequireAndVerifyClientCert,
				ClientCAs:    cPool,
			}
			opts = append(opts,
				grpc.Creds(credentials.NewTLS(&tlsConfig)),
			)
		}
	}

	server := grpc.NewServer(opts...)
	pb.RegisterGreeterServer(server, &Service{})

	logrus.WithField("addr", argAddress).Println("Starting server")
	server.Serve(listener)
}
