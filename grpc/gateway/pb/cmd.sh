protoc -I=. -I=$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:. helloworld.proto
protoc -I=. -I=$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --gogoslick_out=plugins=grpc:. helloworld.proto


go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger

protoc -I=. -I=$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis  --swagger_out=logtostderr=true:. helloworld.proto
protoc -I=. -I=$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis  --grpc-gateway_out=logtostderr=true:. helloworld.proto

