package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/redis.v2"
)

func main() {
	// redisAddr := "127.0.0.1:6379"
	// redisAuth := "yourpassword"
	redisAddr := "10.200.6.99:6379"
	redisAuth := ""

	redisClient := NewRedisClient(redisAddr, redisAuth)

	streamName := "test"
	value := map[string]interface{}{
		"topic":   "test-xadd",
		"content": "something",
	}

	router := gin.Default()
	router.GET("/insert", func(c *gin.Context) {
		id, err := redisClient.xadd(streamName, value)
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
