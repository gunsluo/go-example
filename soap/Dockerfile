FROM golang:1.18.1-buster as gobuilder

RUN apt-get update && apt-get install -y --no-install-recommends build-essential unzip gss-ntlmssp \
	&& rm -rf /var/lib/apt/lists/*

COPY . /build

RUN cd /build && go build -o soap-srv sever.go soap.go && cp -fr soap-srv /usr/local/bin/soap-srv

