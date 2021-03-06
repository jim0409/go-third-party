package service

import (
	"github.com/golang-jwt/jwt"
)

var JwtKey = []byte("my_secret_key")

var Users = map[string]string{
	"user1": "password1",
}

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
