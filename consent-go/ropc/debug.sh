#!/bin/sh

# example
go run main.go \
    -endpoint http://127.0.0.1:4444/ \
    -admin http://127.0.0.1:4445/ \
    -port 8888 \
    -scope openid,offline

