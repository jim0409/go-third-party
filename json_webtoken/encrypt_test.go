package main

import (
	"fmt"
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
)

func decryptJwt(tknStr string) (*Claims, error) {
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	_ = tkn

	if err != nil {
		return nil, err
	}
	return claims, nil
}

func TestEncrypt(t *testing.T) {
	Req := Request{
		Username: "jim",
		Password: "password",
	}

	login := Login{
		Req:    Req,
		Action: "login",
	}
	tknStr := encryptJwt(&login)
	claims, err := decryptJwt(tknStr)
	assert.Nil(t, err)

	fmt.Println(claims.Action)
	fmt.Println(claims.Req.Username)
	fmt.Println(claims.Req.Password)
}
