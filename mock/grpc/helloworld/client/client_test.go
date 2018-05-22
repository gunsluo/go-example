package main

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gunsluo/go-example/grpc/helloworld/pb"
	mockpb "github.com/gunsluo/go-example/grpc/helloworld/pb/mocks"
)

func TestSayHello(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGreeterClient := mockpb.NewMockGreeterClient(ctrl)
	req := &pb.HelloRequest{Name: "luoji"}
	mockGreeterClient.EXPECT().SayHello(
		gomock.Any(),
		req,
	).Return(&pb.HelloReply{Message: "Mocked Interface"}, nil)

	_, err := sayHello(mockGreeterClient, "luoji")
	if err != nil {
		t.Fatalf("unable to sayhello: %s", err)
	}
}
