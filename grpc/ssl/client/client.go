package main

import (
	"context"
	"fmt"

	"github.com/gunsluo/go-example/grpc/ssl/pb"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	sAddress = "0.0.0.0:3264"
	crtFile  = "../cert/server.crt"
)

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
	var opts []grpc.DialOption
	customCred := &customCredential{token: "custom-token"}
	if crtFile != "" {
		fmt.Println("enable credentials in the grpc")
		// target is common name(host name) in the cert file
		creds, err := credentials.NewClientTLSFromFile(crtFile, "target")
		if err != nil {
			panic(err)
		}

		opts = append(opts, grpc.WithTransportCredentials(creds))
		customCred.security = true
	} else {
		opts = append(opts, grpc.WithInsecure())
	}

	// custom credentials
	opts = append(opts, grpc.WithPerRPCCredentials(customCred))

	conn, err := grpc.Dial("127.0.0.1:3264", opts...)
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
