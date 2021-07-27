package main

import (
	"log"

	"github.com/go-redis/redis"
)

func NewRedisClient(addr string, pw string) redisDAO {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pw,
		DB:       0,
	})

	return &redisObject{
		ro: client,
	}
}

type redisObject struct {
	ro *redis.Client
}

type redisDAO interface {
	lrange(string) ([]string, error)
	lpush(string, string) error
}

func (r *redisObject) lrange(key string) ([]string, error) {
	if err := r.ro.Ping().Err(); err != nil {
		log.Println("ping error")
		return nil, err
	}
	return r.ro.LRange(key, 0, -1).Result()
}

func (r *redisObject) lpush(key string, value string) error {
	return r.ro.LPush(key, value).Err()
}
