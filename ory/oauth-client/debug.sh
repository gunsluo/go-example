#!/bin/sh

if [ "$1" = "pkce" ]; then

  if [ "$2" = "2" ]; then
    go run main.go \
      -client-id public-client2 \
      -endpoint http://127.0.0.1:4444 \
      -port 7777 \
      -scope openid,offline \
      -pkce

    exit 0
  fi

  # public-client
  go run main.go \
    -client-id 81ee13a7-c621-43e6-9709-e0a94cbddc33 \
    -endpoint http://127.0.0.1:4444 \
    -port 5555 \
    -scope openid,offline \
    -pkce

  exit 0
fi


# example
# auth-code-client
go run main.go \
    -client-id 6fcba100-46f6-44ec-b15e-075cf712189d \
    -client-secret secret \
    -endpoint http://127.0.0.1:4444 \
    -port 5555 \
    -scope openid,offline

