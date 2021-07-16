package main

import (
	"fmt"

	"gopkg.in/redis.v2"
)

func main() {
	redisAddr := "127.0.0.1:6379"
	redisAuth := "yourpassword"

	redisClient := NewRedisClient(redisAddr, redisAuth)

	streamName := "test"
	value := map[string]interface{}{
		"topic":   "test-xadd",
		"content": "something",
	}

	err := redisClient.xadd(streamName, value)
	if err != nil && err != redis.Nil {
		fmt.Println(err)
	}

	fmt.Println(redisClient.len(streamName))

}
