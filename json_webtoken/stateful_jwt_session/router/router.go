package router

import (
	"net/http"

	"go-third-party/json_webtoken/stateful_jwt_session/service"
)

func init() {
	http.HandleFunc("/signin", service.Signin)
	http.HandleFunc("/welcome", service.Welcome)
}
