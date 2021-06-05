package main

import (
	"log"
	"net/http"

	"go-third-party/redis_related/redigo/session_based/handler"
	"go-third-party/redis_related/redigo/session_based/redispool"
)

func init() {
	redispool.InitCache()
}

func main() {
	http.HandleFunc("/signin", handler.Signin)
	http.HandleFunc("/welcome", handler.Welcome)
	http.HandleFunc("/refresh", handler.Refresh)

	log.Fatal(http.ListenAndServe(":1234", nil))

}
