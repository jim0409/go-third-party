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

	Client.Set("foo", "bar", 0)

	var luaScript = redis.NewScript(`return redis.call("GET" , KEYS[1])`)

	n, err := luaScript.Run(Client, []string{"foo"}).Result()
	if err != nil {
		panic(err)
	}

	fmt.Println(n, err)

}
