protoc -I=. -I=$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:. helloworld.proto
protoc -I=. -I=$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --gogoslick_out=plugins=grpc:. helloworld.proto



mockgen -package mockpb github.com/gunsluo/go-example/grpc/helloworld/pb GreeterClient > mocks/helloworld_mock.go

