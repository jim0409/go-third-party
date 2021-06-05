package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/gomodule/redigo/redis"
)

type RedisMethod interface {
	SetEXKey(string, string) error
	SetNXKey(string, string) error
	GetKey(string) (interface{}, error)
	DelKey(string) error
}

type RedisOperator struct {
	RedisPool *redis.Pool
	c         redis.Conn
}

// set a expire key with default expire time 120s
func (r *RedisOperator) SetEXKey(k, v string) error {
	var err error
	c := r.RedisPool.Get()
	defer c.Close()

	if _, err = c.Do("SET", k, v, "EX", "120"); err != nil {
		return fmt.Errorf("Error happend while SETEX key %v with %v\n", k, v)
	}
	return err
}

func (r *RedisOperator) SetNXKey(k, v string) error {
	var err error
	c := r.RedisPool.Get()
	defer c.Close()

	if _, err = c.Do("SETNX", k, v); err != nil {
		return fmt.Errorf("Error happend while SETNX key %v with %v\n", k, v)
	}
	return err
}

func (r *RedisOperator) DelKey(k string) error {
	var err error
	c := r.RedisPool.Get()
	defer c.Close()

	if _, err = c.Do("DEL", k); err != nil {
		return fmt.Errorf("Error happend while DEL key %v\n", k)
	}
	return err
}

func (r *RedisOperator) GetKey(k string) (interface{}, error) {
	var err error
	var resp interface{}

	c := r.RedisPool.Get()
	defer c.Close()

	if resp, err = c.Do("GET", k); err != nil {
		return nil, fmt.Errorf("Error happend while GET key %v\n", k)
	}
	return resp, err
}

func NewRedisObj() (RedisMethod, error) {
	pool, err := initPool()
	if err != nil {
		return nil, err
	}

	conn, err := pool.Dial()
	if err != nil {
		log.Printf("ERROR: %v", err)
	}
	ping(conn)

	return &RedisOperator{RedisPool: pool}, nil
}

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

// check connection status
func ping(conn redis.Conn) {
	_, err := redis.String(conn.Do("PING"))
	if err != nil {
		panic(fmt.Sprintf("ERROR: fail ping redis conn: %v\n", err))
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

func main() {
	// pool, err := initPool()
	r, err := NewRedisObj()
	if err != nil {
		panic(err)
	}
	r.SetEXKey("jim", "test")
	r.SetNXKey("jim", "test")
	resp, err := r.GetKey("jim")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(resp)
	// fmt.Println(fmt.Sprintf("%v", resp))
	// r.DelKey("jim")
	// resp, err = r.GetKey("jim")
	// fmt.Println(resp)
	// if err != nil {
	// 	log.Println(err)
	// }
}
