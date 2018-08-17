package main

import (
	"context"
	"fmt"

	"gitlab.com/tesgo/kit/proto/ses/pb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:6000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := pb.NewSESClient(conn)

	reply, err := client.Send(context.Background(),
		&pb.SendRequest{
			From:    "gunsluo@gmail.com",
			To:      []string{"gunsluo@gmail.com"},
			Subject: "Amazon SES Test (AWS SDK for Go)",
			Html:    "<html>this is a test</html>",
		})
	if err != nil {
		fmt.Println("unable to send ", err)
	} else {
		fmt.Println("reply:", reply.Id, reply.Status)
	}
}
