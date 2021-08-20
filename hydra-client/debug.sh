#!/bin/sh

if [ "$1" = "pkce" ]; then
  go run main.go \
    -client-id public-client \
    -endpoint http://127.0.0.1:4444/ \
    -port 5555 \
    -scope openid,offline \
    -pkce

    exit 0
fi


# example
go run main.go \
    -client-id auth-code-client \
    -client-secret secret \
    -endpoint http://127.0.0.1:4444/ \
    -port 5555 \
    -scope openid,offline

