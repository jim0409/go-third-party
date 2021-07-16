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

	Client.Set("foo", "bar", -1)

	var luaScript = `return redis.call("INFO")`
	result, err := Client.ScriptLoad(luaScript).Result() //返回的指令碼會產生一個sha1雜湊值,下次用的時候可以直接使用這個值，類似於
	if err != nil {
		panic(err)
	}

	foo := Client.EvalSha(result, []string{})
	fmt.Println(foo.Val())
}
