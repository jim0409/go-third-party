package redispool

import (
	"io"
	"log"
	"os"

	"github.com/gomodule/redigo/redis"
)

var Cache redis.Conn

func InitCache() {
	cachePool, err := initPool()
	if err != nil {
		panic(err)
	}
	Cache = cachePool.Get()
}

func initPool() (*redis.Pool, error) {
	var err error
	var pool *redis.Pool
	pool = &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", "localhost:6379")
			if err != nil {
				log.Printf("ERROR: fail init redis: %s", err.Error())
				os.Exit(1)
			}

			//  set a pre authority check for redis connections
			preCheckRedisAuthority(conn, "yourpassword")
			return conn, err
		},
	}

	// set a retry machanism for redigo
	redigoErrRetry(pool)

	return pool, err
}

// use `AUTH` to have a pre check password step
func preCheckRedisAuthority(conn redis.Conn, pw string) {
	if _, err := conn.Do("AUTH", pw); err != nil {
		conn.Close()
		panic(err)
	}
}

// add an extra key for `ERR`
func redigoErrRetry(pool *redis.Pool) {
	c := pool.Get()
	defer c.Close()
	c.Do("ERR", io.EOF)
	if err := c.Err(); err != nil {
		panic(err)
	}
}
