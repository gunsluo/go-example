package main

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var jwtSignedSecret = []byte("secret")

type statisticClaims struct {
	Type  int `json:"type,omitempty"`
	OrgId int `json:"org_id,omitempty"`
	//Value interface{} `json:"value,omitempty"`
}

// authorization
type authzClaims struct {
	Enable   bool   `json:"enable,omitempty"`
	Resource string `json:"resource,omitempty"`
	Action   string `json:"action,omitempty"`
	Context  string `json:"context,omitempty"`
}

type fsClaims struct {
	Statistic *statisticClaims `json:"statistic,omitempty"`
	Criterion *authzClaims     `json:"authorization,omitempty"`
	jwt.StandardClaims
}

func main() {
	claims := fsClaims{
		Statistic: &statisticClaims{
			Type:  1,
			OrgId: 100,
		},
		Criterion: &authzClaims{
			Enable:   true,
			Resource: "aaa",
			Action:   "post",
		},
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
			Issuer:    "service, such as profile",
			Audience:  "fs",
			IssuedAt:  time.Now().Unix(),
			NotBefore: time.Now().Unix(),
			Subject:   "luoji",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := token.SignedString(jwtSignedSecret)

	fmt.Println("token:", signedToken)
}
