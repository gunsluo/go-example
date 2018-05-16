package main

import (
	"context"
	"fmt"

	oidc "github.com/coreos/go-oidc"
)

type AuthInfo struct {
	AccessToken string
	Subject     string
	IdToken     *oidc.IDToken
	UserInfo    *oidc.UserInfo
}

var AuthInfoCtx = struct{}{}

func main() {
	ctx := context.Background()

	authInfo := AuthInfo{AccessToken: "aaa"}

	ctx = context.WithValue(ctx, AuthInfoCtx, authInfo)
	step(ctx)
}

func step(ctx context.Context) {
	v := ctx.Value(AuthInfoCtx)
	if v == nil {
		return
	}

	authInfo, ok := v.(AuthInfo)
	fmt.Println("-->", authInfo, ok)
}
