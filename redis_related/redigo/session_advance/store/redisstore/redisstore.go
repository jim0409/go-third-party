package redisstore

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-redis/redis"
	"go-third-party/redis_related/redigo/session_advance/store"
)

type RedisStore struct {
	client *redis.Client
}

func NewRedisStore() store.Store {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "yourpassword",
		DB:       0, //use default DB
	})
	if _, err := client.Ping().Result(); err != nil {
		log.Fatalf("Failed to ping Redis: %v", err)
	}
	return &RedisStore{
		client: client,
	}
}

func (r *RedisStore) Set(id string, session store.Session) error {
	bs, err := json.Marshal(session)
	if err != nil {
		return fmt.Errorf("An error occur while unmarshal %v\n", err)
	}

	if err := r.client.Set(id, bs, 0).Err(); err != nil {
		return fmt.Errorf("An error occur while set into redis %v\n", err)
	}

	return nil
}

func (r *RedisStore) Get(id string) (store.Session, error) {
	var session store.Session

	bs, err := r.client.Get(id).Bytes()
	if err != nil {
		return session, fmt.Errorf("An error occur while get session info %v\n", err)
	}

	if err := json.Unmarshal(bs, &session); err != nil {
		return session, fmt.Errorf("An error occure while unmarsh info %v\n", err)
	}

	return session, nil
}
