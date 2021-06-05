package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/gomodule/redigo/redis"
)

// initial redis connection pool
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

// check connection status
func ping(conn redis.Conn) {
	_, err := redis.String(conn.Do("PING"))
	if err != nil {
		panic(fmt.Sprintf("ERROR: fail ping redis conn: %v\n", err))
	}
}

func main() {
	pool, err := initPool()
	conn, err := pool.Dial()
	if err != nil {
		log.Printf("ERROR: %v", err)
	}
	ping(conn)

}

// func set
