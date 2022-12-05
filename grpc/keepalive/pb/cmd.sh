
export PATH=$PATH:/Users/luoji/gopath/src/tespkg.in/meera-proto/.bin
#protoc -I=. --go_out=plugins=grpc:. helloworld.proto
protoc --go_out=. --go-grpc_out=. helloworld.proto
mv github.com/gunsluo/go-example/grpc/keepalive/pb/*.go .
rm -fr github.com



