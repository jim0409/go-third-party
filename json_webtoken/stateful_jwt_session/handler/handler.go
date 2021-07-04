package handler

import (
	"net/http"

	"github.com/jimweng/json_webtoken/stateful_jwt_session/handler/signin"
	"github.com/jimweng/json_webtoken/stateful_jwt_session/handler/welcome"
)

func init() {
	http.HandleFunc("/signin", signin.Signin)
	http.HandleFunc("/welcome", welcome.Welcome)
}
