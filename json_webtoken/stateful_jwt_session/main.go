package main

import (
	"log"
	"net/http"

	_ "github.com/jimweng/json_webtoken/stateful_jwt_session/handler"
	"github.com/jimweng/json_webtoken/stateful_jwt_session/redispool"
)

func main() {
	log.Fatal(http.ListenAndServe(":8000", nil))

}

func init() {
	redispool.InitCache()
}
