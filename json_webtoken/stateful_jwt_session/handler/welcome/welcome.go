package welcome

import (
	"fmt"
	"go-third-party/json_webtoken/stateful_jwt_session/handler/conf"
	"go-third-party/json_webtoken/stateful_jwt_session/redispool"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	// We can obtain the session token from the requests cookies, which come with every request
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// For any other type of error, return a bad request status
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	sessionToken := c.Value
	response, err := redispool.Cache.Do("GET", sessionToken)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if response == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Get the JWT string from the cookie
	tknStr := fmt.Sprintf("%s", response)

	// remove the used token ... if it is a one consume service
	if _, err := redispool.Cache.Do("DEL", sessionToken); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	// Initialize a new instance of `Claims`
	claims := &conf.Claims{}

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return conf.JwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Finally, return the welcome message to the user, along with their
	// username given in the token
	w.Write([]byte(fmt.Sprintf("Welcome %s!", claims.Username)))

}
