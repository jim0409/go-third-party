package main

import "github.com/go-redis/redis"

type RedisInstance struct {
	// Client *redis.ClusterClient
	Client *redis.Client
}

type IRedisMethod interface {
	HSet(key, field string, value int) error
	HIncr(key, field string, value int) error

	Set(key string, value int) error
	Incr(key string, value int) error

	Lpush(key string, value int) error
}

// Cluster
func NewRedisInstance() IRedisMethod {
	// clusterClient := redis.NewClusterClient(&redis.ClusterOptions{
	// 	Addrs: []string{
	// 		":7001",
	// 		":7002",
	// 		":7003",
	// 		":7004",
	// 		":7005",
	// 		":7006",
	// 	},
	// 	PoolSize: 1000,
	// })
	// return &RedisInstance{
	// 	Client: clusterClient,
	// }

	client := redis.NewClient(&redis.Options{
		Addr:     ":6379",
		Password: "yourpassword",
		PoolSize: 1000,
	})
	return &RedisInstance{
		Client: client,
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
	return r.Client.LPush(key, value).Err()
}
