package main

import (
	"fmt"
	"sync"

	"github.com/go-redis/redis"
)

var (
	redisDbConfig   *RedisConfig
	redisPoolObject RedisDBAccessObject
	once            sync.Once
)

type RedisConfig struct {
	DBAddr string
	DBAuth string
}

type redisPoolObj struct {
	ro *redis.Client
}

type RedisDBAccessObject interface {
	SetKey(string, interface{}) error
	// DeleteKey(string) error
	LuaScript(string)
	Close() error
}

func LoadRedisDBConfig(dbAddr, dbAuth string) {
	redisDbConfig = &RedisConfig{
		DBAddr: dbAddr,
		DBAuth: dbAuth,
	}
}

func RetriveRedisPoolObj() RedisDBAccessObject {
	once.Do(func() {
		// redisPoolObject = redisPoolObject
		redisPoolObject = &redisPoolObj{}
	})
	return redisPoolObject
}

func StartRedisPool() error {
	var err error
	redisPoolObject = RetriveRedisPoolObj()
	redisPoolObject, err = initRedisDB(redisDbConfig)
	return err
}

func initRedisDB(r *RedisConfig) (RedisDBAccessObject, error) {
	fmt.Println("start redis client")
	client := redis.NewClient(&redis.Options{
		Addr:     r.DBAddr,
		Password: r.DBAuth,
	})

	return &redisPoolObj{
		ro: client,
	}, nil

}

// ====== redisPoolObj method ===
func (rdb *redisPoolObj) SetKey(key string, value interface{}) error {
	return rdb.ro.Set(key, value, 0).Err()
}

func (rdb *redisPoolObj) Close() error {
	return rdb.ro.Close()
}

func (rdb *redisPoolObj) LuaScript(s string) {
	// return redis.call("GET" , KEYS[1])
	luaScript := redis.NewScript(s)
	// n, err := luaScript.Run(rdb.ro, []string{"jim"}).Result()
	n, err := luaScript.Run(rdb.ro, []string{"jim"}).Result()
	fmt.Println(n, err)
}

// ====== redisPoolObj method ===

// ====== redisScript ======

func main() {
	RetriveRedisPoolObj()
	LoadRedisDBConfig("127.0.0.1:6379", "yourpassword")
	// initRedisDB(redisDbConfig)
	if err := StartRedisPool(); err != nil {
		panic(err)
	}

	// localRedisObj := RetriveRedisPoolObj()
	redisPoolObject.SetKey("jim", 123)

	luaScript := `
	local elements = redis.call("GET", KEYS[1])
	return elements+1
	`
	// return redis.call("GET" , KEYS[1])

	redisPoolObject.LuaScript(luaScript)
}
