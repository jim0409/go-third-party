package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis"
)

func NewRedisClient(addr string, pw string) redisDAO {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pw,
		DB:       0,
	})

	_, err := client.Do("CLIENT", "SETNAME", "demo").Result()
	if err != nil {
		panic(err)
	}

	rO := &redisObject{
		ro: client,
	}
	return rO
}

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s <redis-addr> <redis-auth>\n",
			os.Args[0])
		os.Exit(1)
	}
	redisAddr := os.Args[1] // 127.0.0.1:6379
	redisAuth := os.Args[2] // yourpassword

	// redisClient := NewRedisClient("localhost:6379", "yourpassword")
	redisClient := NewRedisClient(redisAddr, redisAuth)
	fmt.Println(redisClient.ping())
	// fmt.Printf("%v\n", pong)
	fmt.Println(redisClient.setValue("a", "jim"))
	fmt.Println(redisClient.queryValue("a"))
	fmt.Println(redisClient.deleteValue("a"))
	// lpush lpush_key lpush_value
	// fmt.Println(redisClient.LPOP("lpush_key"))

	for {
		time.Sleep(1 * time.Second)
		fmt.Println(redisClient.ping())
	}
}

type redisObject struct {
	ro *redis.Client
}

type redisDAO interface {
	ping() string                  // expect to print out pong on the console
	setValue(string, string) error // expect to set value into redis
	queryValue(string) string
	deleteValue(string) error
	LPOP(string) string
}

func (r *redisObject) LPOP(s string) string {
	v, err := r.ro.LPop(s).Result()
	if err != nil {
		log.Panic(err)
	}
	return v
}

func (r *redisObject) ping() string {
	pong, _ := r.ro.Ping().Result()
	return pong
}

func (r *redisObject) setValue(key string, value string) error {
	err := r.ro.Set(key, value, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *redisObject) queryValue(key string) string {
	if val, err := r.ro.Get(key).Result(); err != nil {
		return fmt.Sprintf("Error happend with %v\n", err)
	} else {
		return fmt.Sprintf("Query key '%s', get return value '%s'\n", key, val)
	}
}

func (r *redisObject) deleteValue(key string) error {
	_, err := r.ro.Del(key).Result()
	return err
}

// if err == redis.Nil {
// 	fmt.Println("key2 does not exist")
// } else if err != nil {
// 	panic(err)
// } else {
// 	fmt.Println("key2", val2)
// }
// Output: key value
// key2 does not exist
