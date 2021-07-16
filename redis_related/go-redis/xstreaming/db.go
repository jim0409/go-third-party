package main

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis"
)

func NewRedisClient(addr string, pw string) redisDAO {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pw,
		DB:       0,
	})
	rO := &redisObject{
		ro: client,
	}
	return rO
}

type redisObject struct {
	ro *redis.Client
}

type redisDAO interface {
	len(string) int64
	xread(string)
	xadd(string, map[string]interface{}) error
}

func (r *redisObject) len(streamName string) int64 {
	l, err := r.ro.XLen(streamName).Result()
	if err != nil {
		log.Printf("retrive stream len failed %v\n", streamName)
		return 0
	}
	return l
}

func (r *redisObject) xadd(streamName string, value map[string]interface{}) error {
	tmp := r.len(streamName)

	err := r.ro.XAdd(&redis.XAddArgs{
		Stream: streamName,
		Values: value,
	},
	).Err()
	if err != nil {
		return err
	}

	if tmp == r.len(streamName) {
		return fmt.Errorf("write_failed")
	}

	return nil
}

func (r *redisObject) xread(streamName string) {
	entries, err := r.ro.XRead(&redis.XReadArgs{
		Streams: []string{streamName, "0-1000"},
		Count:   1,
		Block:   2 * time.Millisecond,
	}).Result()

	if err != nil {
		if errors.Is(err, redis.Nil) {
			return
		}
		log.Fatal(err)
		return
	}

	for _, msg := range entries[0].Messages {
		fmt.Println(msg)
		fmt.Println(r.ro.XDel(streamName, msg.ID).Result())
	}
}
