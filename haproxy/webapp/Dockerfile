FROM golang:1.13-buster as builder

ENV GO111MODULE off 

COPY . /go/src/webapp/
RUN cd /go/src/webapp && go build -o room-srv main.go

FROM debian:buster-slim

RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates libaio1 && rm -rf /var/lib/apt/lists/*

COPY --from=builder /go/src/webapp/room-srv /usr/local/bin/

EXPOSE 54321

WORKDIR "/usr/local/bin"
