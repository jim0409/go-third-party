package redisdb

import (
	"errors"
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/gomodule/redigo/redis"
)

var (
	redisDbConfig   *RedisConfig
	redisPoolObject RedisDBAccessObject
	once            sync.Once
)

type RedisConfig struct {
	DBAddr    string
	DBAuth    string
	MaxIdle   int
	MaxActive int
	TimeOut   int
}

func LoadRedisDBConfig(dbAddr, dbAuth string, dbMaxIdle, dbMaxActive, timeOut int) {
	redisDbConfig = &RedisConfig{
		DBAddr:    dbAddr,
		DBAuth:    dbAuth,
		MaxIdle:   dbMaxIdle,
		MaxActive: dbMaxActive,
		TimeOut:   timeOut,
	}
}

func RetriveRedisPoolObj() RedisDBAccessObject {
	once.Do(func() {
		// redisPoolObject = redisPoolObject
		redisPoolObject = &redisPoolObj{}
	})
	return redisPoolObject
}

type redisPoolObj struct {
	Pool *redis.Pool
}

func (rdb *redisPoolObj) GetKeyValue(key string) (interface{}, error) {
	c := rdb.Pool.Get()
	defer c.Close()

	value, err := redis.String(c.Do("GET", key))
	switch err {
	case nil:
		return value, nil

	case redis.ErrNil:
		return "", errors.New("key not exist")

	default:
		return "", err

	}
}

func (rdb *redisPoolObj) SetKey(key string, value interface{}) error {
	c := rdb.Pool.Get()
	defer c.Close()

	if _, err := c.Do("SET", key, value); err != nil {
		return err
	}

	return nil
}

func (rdb *redisPoolObj) SetExpiredKey(key string, value interface{}, exTime int) error {
	c := rdb.Pool.Get()
	defer c.Close()

	if _, err := c.Do("SET", key, value, "EX", exTime); err != nil {
		return err
	}

	return nil
}

func (rdb *redisPoolObj) SetNonExistedKey(key string, value interface{}) error {
	c := rdb.Pool.Get()
	defer c.Close()

	// if meet error or unavailable to set replicated key return errors
	if result, err := c.Do("SETNX", key, value); err != nil || fmt.Sprintf("%v", result) != "1" {
		return errors.New("Failed to set non-existed key to redis")
	}

	return nil
}

func (rdb *redisPoolObj) DeleteKey(key string) error {
	c := rdb.Pool.Get()
	defer c.Close()

	if _, err := c.Do("DEL", key); err != nil {
		return err
	}

	return nil
}

func (rdb *redisPoolObj) Close() error {
	return rdb.Close()
}

type RedisDBAccessObject interface {
	// CheckKeyExistOrNot(string) (bool, error) // .. not sure need it or not
	GetKeyValue(string) (interface{}, error)
	SetKey(string, interface{}) error
	SetExpiredKey(string, interface{}, int) error
	SetNonExistedKey(string, interface{}) error
	DeleteKey(string) error
	Close() error
}

func StartRedisPool() error {
	var err error
	redisPoolObject = RetriveRedisPoolObj()
	redisPoolObject, err = initRedisDB(redisDbConfig)
	return err
}

// InitMasterRedis init master redis
func initRedisDB(r *RedisConfig) (RedisDBAccessObject, error) {
	proto := "tcp"

	pool := &redis.Pool{
		MaxIdle:     r.MaxIdle,
		MaxActive:   r.MaxActive,
		IdleTimeout: time.Duration(r.TimeOut) * time.Second,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial(proto, r.DBAddr)
			if err != nil {
				return nil, fmt.Errorf("redis proto=%v addr=%v dial err: %v", proto, r.DBAddr, err)
			}

			if _, err = conn.Do("AUTH", r.DBAuth); err != nil {
				conn.Close()
				return nil, fmt.Errorf("redis set addr=%v auth=%v is err: %v", r.DBAddr, r.DBAuth, err)
			}
			return conn, nil
		},
	}

	c := pool.Get()
	defer c.Close()
	c.Do("ERR", io.EOF)
	if c.Err() != nil {
		return nil, fmt.Errorf("redis do err: %v", c.Err())
	}

	return &redisPoolObj{
		Pool: pool,
	}, nil
}
