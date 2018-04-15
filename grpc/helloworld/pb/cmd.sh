protoc -I=. -I=$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:. helloworld.proto
protoc -I=. -I=$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --gogoslick_out=plugins=grpc:. helloworld.proto

