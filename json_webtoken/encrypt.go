package main

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("test")

type Login struct {
	Req    Request `json:"Request"`
	Action string  `json:"Action"`
}

type Request struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Login
	jwt.StandardClaims
}

func encryptJwt(l *Login) string {
	claims := &Claims{
		Login:          *l,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			// ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		panic(err)
	}
	return tokenString
}

func main() {
	Req := Request{
		Username: "jim",
		Password: "password",
	}

	login := Login{
		Req:    Req,
		Action: "login",
	}

	fmt.Println(encryptJwt(&login))
}
