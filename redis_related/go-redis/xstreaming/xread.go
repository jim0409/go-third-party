package main

import "time"

func main() {
	// redisAddr := "127.0.0.1:6379"
	// redisAuth := "yourpassword"
	redisAddr := "10.200.6.99:6379"
	redisAuth := ""

	redisClient := NewRedisClient(redisAddr, redisAuth)

	streamName := "test"
	consumerGroup := "testgp"

	if err := redisClient.InitXGroup(streamName, consumerGroup); err != nil {
		panic(err)
	}

	// go cron(redisClient, streamName, consumerGroup, "jim")
	for {
		redisClient.XReadGroup(streamName, consumerGroup, "jim")
	}

}

func cron(r redisDAO, streamName string, consumerGroup string, consumerName string) {
	ticker := time.Tick(time.Second * 3)
	for {
		select {
		case <-ticker:
			r.ConsumePendingGroup(streamName, consumerGroup, consumerName)
		}
	}
}
