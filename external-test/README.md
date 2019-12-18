go test -coverprofile=coverage.out -c main.go main_test.go -o test

./test -test.coverprofile=coverage.out -test.v -test.run=TestMain

go test -coverprofile=coverage.out -run=TestMain

curl -XPOST http://127.0.0.1:1234/test
curl -XPOST http://127.0.0.1:1234/deathblow

go tool cover -func=coverage.out


