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
    -client-id 968c1afc-335f-4373-83cc-5e36bdd98896 \
    -endpoint http://127.0.0.1:4444 \
    -port 5555 \
    -scope openid,offline \
    -pkce

  exit 0
fi


# example
# auth-code-client
go run main.go \
    -client-id a8ea5381-7e4b-4558-ba3f-4a6ccc9e2d5d \
    -client-secret secret \
    -endpoint http://127.0.0.1:4444 \
    -port 5555 \
    -scope openid,offline

