package handler

import (
	"fmt"
	"net/http"
	"time"

	"go-third-party/redis_related/redigo/session_based/redispool"
)

func Refresh(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
	}
	sessionToken := c.Value

	response, err := redispool.Cache.Do("GET", sessionToken)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if response == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	newSessionToken := ReturnUUidNewV4()
	_, err = redispool.Cache.Do("SETEX", newSessionToken, "120", fmt.Sprintf("%s", response))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = redispool.Cache.Do("DEL", sessionToken)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Set the new token as the users `session_token` cookie
	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   newSessionToken,
		Expires: time.Now().Add(120 * time.Second),
	})
}
