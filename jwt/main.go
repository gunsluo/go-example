package main

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("secret")

type user struct {
	UserID   uint64 `json:"userID"`
	UserName string `json:"userName"`
	IsAdmin  bool   `json:"isAdmin"`
}

type userClaims struct {
	user
	jwt.StandardClaims
}

func main() {
	token, err := login("admin", "admin")
	if err != nil {
		panic(err)
	}
	fmt.Println("login token:", token)

	if ok := auth(token); ok {
		fmt.Println("auth pass")
	} else {
		fmt.Println("auth failed")
	}
}

func login(username, password string) (string, error) {
	userParam := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{
		Username: username,
		Password: password,
	}

	if userParam.Username != "admin" || userParam.Password != "admin" {
		return "", fmt.Errorf("invalid login")
	}

	//generate token
	expire := time.Now().Add(time.Hour * 1).Unix()
	//var expire int64 = 1518013920
	// Create the Claims
	claims := userClaims{
		user: user{
			UserID:   1,
			UserName: userParam.Username,
			IsAdmin:  true,
		},
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expire,
			Issuer:    "login",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := token.SignedString(jwtSecret)

	return signedToken, nil
}

func auth(jwtToken string) bool {
	// parse tokentoken
	token, err := jwt.ParseWithClaims(jwtToken, &userClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return jwtSecret, nil
	})
	if err != nil {
		return false
	}

	_, ok := token.Claims.(*userClaims)
	if !ok || !token.Valid {
		return false
	}

	return true
}
