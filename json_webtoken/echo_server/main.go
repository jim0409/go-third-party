package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
)

// Create the JWT key used to create the signature
var jwtKey = []byte(os.Getenv("JWT_KEY"))

type Login struct {
	Request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"Request"`
	Action string `json:"Action"`
}

type Claims struct {
	Login
	jwt.StandardClaims
}

func main() {
	// "Signin" and "Welcome" are the handlers that we will implement
	http.HandleFunc("/echo", EchoEncrypt)

	// start the server on port 8000
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func EchoEncrypt(w http.ResponseWriter, r *http.Request) {
	var lg Login

	if err := json.NewDecoder(r.Body).Decode(&lg); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// expirationTime := time.Now().Add(5 * time.Minute)

	claims := &Claims{
		Login:          lg,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			// ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	io.WriteString(w, tokenString)
}
