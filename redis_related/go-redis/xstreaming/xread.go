package main

import "fmt"

func main() {
	redisAddr := "127.0.0.1:6379"
	redisAuth := "yourpassword"

	redisClient := NewRedisClient(redisAddr, redisAuth)

	streamName := "test"

	fmt.Println(redisClient.len(streamName))
	redisClient.xread(streamName)
}
