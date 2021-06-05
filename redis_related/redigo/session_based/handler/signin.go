package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"go-third-party/redis_related/redigo/session_based/redispool"
)

// 處理登入的流程
func Signin(w http.ResponseWriter, r *http.Request) {
	var creds Credentials

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	expectedPassword, ok := users[creds.Username]
	if !ok || expectedPassword != creds.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	sessionToken := ReturnUUidNewV4()
	_, err = redispool.Cache.Do("SETEX", sessionToken, "120", creds.Username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   sessionToken,
		Expires: time.Now().Add(123 * time.Second),
	})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(fmt.Sprintf("{'sessionToken': '%v'}", sessionToken))

}
