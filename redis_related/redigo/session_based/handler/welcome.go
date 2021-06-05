package handler

import (
	"fmt"
	"net/http"

	"go-third-party/redis_related/redigo/session_based/redispool"
)

// 處理顯示使用者資訊的流程
func Welcome(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
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
	w.Write([]byte(fmt.Sprintf("Welcome %s!", response)))
}
