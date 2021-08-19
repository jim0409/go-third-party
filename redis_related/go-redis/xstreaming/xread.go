package main

func main() {
	// redisAddr := "127.0.0.1:6379"
	redisAddr := "10.200.6.99:6379"
	// redisAuth := "yourpassword"
	redisAuth := ""

	redisClient := NewRedisClient(redisAddr, redisAuth)

	streamName := "test"
	consumerGroup := "testgp"

	// for {
	// 	if redisClient.len(streamName) != 0 {
	// 		redisClient.xread(streamName)
	// 	}
	// }

	if err := redisClient.InitXGroup(streamName, consumerGroup); err != nil {
		panic(err)
	}
	for {
		redisClient.xGroupRead(streamName, consumerGroup, "jim")
	}
}
