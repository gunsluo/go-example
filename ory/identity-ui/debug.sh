#!/bin/sh

if [ "$1" = "swagger" ]; then
  export SWG_PATH=/Users/luoji/gopath/src/xbpkg.io/identity/spec/api.json

  openapi-generator generate -i ${SWG_PATH} -g go -o ./swagger/identityclient \
    --git-user-id gunsluo \
    --git-repo-id "go-example/ory/identity-ui/identityclient" \
    --git-host github.com \
    -t swagger/templates/go \
    -c swagger/go.yml

  rm -fr ./swagger/identityclient/go.mod ./swagger/identityclient/go.sum

  exit 0
fi

export IDENTITY_ENDPOINT=http://127.0.0.1:4544
export PORT=4555
go run main.go
