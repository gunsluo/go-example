package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gunsluo/go-example/grpc/ssl/pb"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	argGRPCAddr   string
	argCrtFile    string
	argKeyFile    string
	argCAFile     string
	argCNOverride string
)

var verbose bool
var rootCmd *cobra.Command

func init() {
	rootCmd = &cobra.Command{
		Use:   "grpc-client",
		Short: "demo client",
		Long:  "Top level command for demo client",
		Run:   run,
	}

	rootCmd.Flags().StringVarP(&argGRPCAddr, "grpc-addr", "a", "127.0.0.1:3264", "grpc address")
	rootCmd.Flags().StringVar(&argCrtFile, "cert-file", "", "certificate file for gRPC TLS authentication")
	rootCmd.Flags().StringVar(&argKeyFile, "key-file", "", "key file for gRPC TLS authentication")
	rootCmd.Flags().StringVar(&argCAFile, "ca-file", "", "ca file for gRPC client")
	rootCmd.Flags().StringVar(&argCNOverride, "cn-override", "", "domain name override")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
}

// customCredential 自定义认证
type customCredential struct {
	token    string
	security bool
}

func (c customCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"token": c.token,
	}, nil
}

func (c customCredential) RequireTransportSecurity() bool {
	return c.security
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run(cmd *cobra.Command, _ []string) {
	var opts []grpc.DialOption
	customCred := &customCredential{token: "custom-token"}
	if argCAFile != "" {
		fmt.Println("enable credentials in the grpc")
		if argCrtFile != "" && argKeyFile != "" {
			cPool := x509.NewCertPool()
			caCert, err := ioutil.ReadFile(argCAFile)
			if err != nil {
				panic(err)
				//return nil, fmt.Errorf("invalid CA crt file: %s", caPath)
			}
			if cPool.AppendCertsFromPEM(caCert) != true {
				panic(err)
				//return nil, fmt.Errorf("failed to parse CA crt")
			}

			clientCert, err := tls.LoadX509KeyPair(argCrtFile, argKeyFile)
			if err != nil {
				panic(err)
				//return nil, fmt.Errorf("invalid client crt file: %s", caPath)
			}

			clientTLSConfig := &tls.Config{
				RootCAs:      cPool,
				Certificates: []tls.Certificate{clientCert},
			}
			cred := credentials.NewTLS(clientTLSConfig)

			opts = append(opts, grpc.WithTransportCredentials(creds))
		} else {

			// target is common name(host name) in the cert file
			creds, err := credentials.NewClientTLSFromFile(argCAFile, argCNOverride)
			if err != nil {
				panic(err)
			}

			opts = append(opts, grpc.WithTransportCredentials(creds))
		}

		customCred.security = true
	} else {
		opts = append(opts, grpc.WithInsecure())
	}

	// custom credentials
	opts = append(opts, grpc.WithPerRPCCredentials(customCred))

	conn, err := grpc.Dial(argGRPCAddr, opts...)
	if err != nil {
		panic(err)
	}
	greeterClient := pb.NewGreeterClient(conn)

	reply, err := greeterClient.SayHello(context.Background(),
		&pb.HelloRequest{Name: "luoji"})
	if err != nil {
		logrus.WithError(err).Fatal("unable to sayhello")
	}

	logrus.Info("reply:", reply.Message)
	conn.Close()
}
