package main

import (
	"go-third-party/json_webtoken/stateful_jwt_session/redispool"
	_ "go-third-party/json_webtoken/stateful_jwt_session/router"
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":8000", nil))

}

func init() {
	redispool.InitCache()
}
