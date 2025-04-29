package main

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	hmacSampleSecret = []byte("secret")
)

func main() {

	// Create claims with multiple fields populated
	claims := MyCustomClaims{
		"bar",
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "test",
			Subject:   "somebody",
			ID:        "1",
			Audience:  []string{"somebody_else"},
		},
	}

	fmt.Printf("foo: %v\n", claims.Foo)

	// Create claims while leaving out some of the optional fields
	// claims = MyCustomClaims{
	// 	"bar",
	// 	jwt.RegisteredClaims{
	// 		// Also fixed dates can be used for the NumericDate
	// 		ExpiresAt: jwt.NewNumericDate(time.Unix(1516239022, 0)),
	// 		Issuer:    "test",
	// 	},
	// }

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(hmacSampleSecret)
	fmt.Println(tokenString, err)

	token, err = jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return hmacSampleSecret, nil
	}, jwt.WithLeeway(5*time.Second))
	if err != nil {
		log.Fatal(err)
	} else if claims, ok := token.Claims.(*MyCustomClaims); ok {
		fmt.Println(claims.Foo, claims.RegisteredClaims.Issuer)
	} else {
		log.Fatal("unknown claims type, cannot proceed")
	}

	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	// 	"foo": "bar",
	// 	"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	// })
	//
	// // Sign and get the complete encoded token as a string using the secret
	// tokenString, err := token.SignedString(hmacSampleSecret)
	//
	// fmt.Println(tokenString, err)
}

type MyCustomClaims struct {
	Foo string `json:"foo"`
	jwt.RegisteredClaims
}
