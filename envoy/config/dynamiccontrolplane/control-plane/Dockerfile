FROM golang:1.15.6-buster as builder

COPY . /go/src/cp/
RUN cd /go/src/cp && go build -o cp-srv main.go

FROM debian:buster-slim

RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates libaio1 && rm -rf /var/lib/apt/lists/*

COPY --from=builder /go/src/cp/cp-srv /usr/local/bin/

WORKDIR "/usr/local/bin"
