FROM golang:1.11.0-alpine3.8 as builder

RUN apk add -u --no-cache git

COPY . /go/src/github.com/gunsluo/go-example/grpc/ssl
RUN go build -o /go/bin/grpc-serve /go/src/github.com/gunsluo/go-example/grpc/ssl/server

FROM alpine

ENV "GOPATH" "/go"

RUN apk add -u --no-cache tzdata ca-certificates && rm -rf /var/cache/apk/*

COPY --from=builder /go/bin/grpc-serve /usr/local/bin/

WORKDIR "/"
