package handler

import (
	"go-third-party/json_webtoken/stateful_jwt_session/handler/signin"
	"go-third-party/json_webtoken/stateful_jwt_session/handler/welcome"
	"net/http"
)

func init() {
	http.HandleFunc("/signin", signin.Signin)
	http.HandleFunc("/welcome", welcome.Welcome)
}
