package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

var Client *redis.Client

func init() {
	Client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

//noinspection GoInvalidCompositeLiteral
func main() {

	Client.FlushAll()

	foo := []redis.Z{
		{
			1732, "George Washington",
		},
		{
			1809, "Abraham Lincoln",
		},
		{
			1858, "Theodore Roosevelt",
		},
	}

	Client.ZAdd("presidents", foo...)

	var luaScript = redis.NewScript(`
        local elements = redis.call("ZRANGE" , KEYS[1] , 0 , 0) 
        redis.call("ZREM" , KEYS[1] , elements[1])
        return elements[1]
    `)

	n, err := luaScript.Run(Client, []string{"presidents"}, 1).Result()
	if err != nil {
		panic(err)
	}

	fmt.Println(n, err)

}
