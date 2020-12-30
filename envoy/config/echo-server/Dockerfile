FROM golang:1.15.6-buster as builder

ENV GO111MODULE off 

COPY . /go/src/echo/
RUN cd /go/src/echo && go build -o echo-srv main.go

FROM debian:buster-slim

RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates libaio1 && rm -rf /var/lib/apt/lists/*

COPY --from=builder /go/src/echo/echo-srv /usr/local/bin/

WORKDIR "/usr/local/bin"
