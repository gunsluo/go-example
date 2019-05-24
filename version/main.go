package main

import "fmt"

var (
	VERSION    string
	BUILD_TIME string
	GO_VERSION string
)

func main() {
	// go build -ldflags "-X main.VERSION=1.0.0 -X 'main.BUILD_TIME=`date`' -X 'main.GO_VERSION=`go version`'"
	fmt.Printf("%s\n%s\n%s\n", VERSION, BUILD_TIME, GO_VERSION)
}
