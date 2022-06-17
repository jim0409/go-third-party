package main

import (
	"github.com/go-redis/redis"
)

type RedisInstance struct {
	Client redis.UniversalClient
}

// Cluster
func NewRedisInstance() *RedisInstance {
	opt := redis.UniversalOptions{
		Addrs: []string{
			// ":6379",
			":7001",
			":7002",
			":7003",
			":7004",
			":7005",
			":7006",
		},
		// Password: "yourpassword",
		PoolSize: 10000,
	}
	clusterClient := redis.NewUniversalClient(&opt)
	_, err := clusterClient.Ping().Result()
	if err != nil {
		panic(err)
	}

	return &RedisInstance{
		Client: clusterClient,
	}
}

func (r *RedisInstance) HSet(key, field string, value int) error {
	return r.Client.HSet(key, field, value).Err()
}

func (r *RedisInstance) HIncr(key, field string, value int) error {
	return r.Client.HIncrBy(key, field, int64(value)).Err()
}

func (r *RedisInstance) Set(key string, value int) error {
	return r.Client.Set(key, int64(value), 0).Err()
}

func (r *RedisInstance) Incr(key string, value int) error {
	return r.Client.IncrBy(key, int64(value)).Err()
}

func (r *RedisInstance) Lpush(key string, value int) error {
	pipe := r.Client.Pipeline()
	pipe.LPush(key, value)
	defer pipe.Exec()
	return nil
}
