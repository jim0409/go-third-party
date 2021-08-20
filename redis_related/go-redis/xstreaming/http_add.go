package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gopkg.in/redis.v2"
)

var (
	// redisAddr := "127.0.0.1:6379"
	// redisAuth := "yourpassword"
	redisAddr = "10.200.6.99:6379"
	redisAuth = ""
)

func main() {
	if addr := os.Getenv("REDIS_ADDR"); addr != "" {
		redisAddr = addr
	}

	if auth := os.Getenv("REDIS_AUTH"); auth != "" {
		redisAuth = auth
	}

	redisClient := NewRedisClient(redisAddr, redisAuth)

	streamName, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	value := map[string]interface{}{
		"topic":   "test-xadd",
		"content": "something",
	}

	router := gin.New()
	router.Use(gin.Recovery())
	router.GET("/insert", func(c *gin.Context) {
		id, err := redisClient.XAdd(streamName, value)
		if err != nil && err != redis.Nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "error",
				"msg":    err,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"id":     id,
		})
	})

	httpSrv := &http.Server{
		Addr:    ":8001",
		Handler: router,
	}

	if err := httpSrv.ListenAndServe(); err != nil {
		panic(err)
	}

}
