package main

import (
	"errors"
	"fmt"
	"log"
	"strings"
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
	xadd(string, map[string]interface{}) (string, error)
	xGroupRead(string, string, string)
	InitXGroup(string, string) error
}

func (r *redisObject) len(streamName string) int64 {
	l, err := r.ro.XLen(streamName).Result()
	if err != nil {
		log.Printf("retrive stream len failed %v\n", streamName)
		return 0
	}
	return l
}

func (r *redisObject) xadd(streamName string, value map[string]interface{}) (string, error) {
	newId, err := r.ro.XAdd(&redis.XAddArgs{
		Stream: streamName,
		Values: value,
	}).Result()
	if err != nil {
		return "", err
	}

	return newId, nil
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

// xread 與 xgroupread 差異:
// http://www.redis.cn/commands/xreadgroup.html
func (r *redisObject) xGroupRead(streamName string, groupName string, consumerName string) {
	start := ">"
	entries, err := r.ro.XReadGroup(&redis.XReadGroupArgs{
		Streams:  []string{streamName, start},
		Group:    groupName,
		Consumer: consumerName,
		Count:    10,
		Block:    0,
	}).Result()

	if err != nil {
		log.Printf("err on consume events: %+v\n", err)
		return
	}

	for _, entry := range entries[0].Messages {
		// waitGrp.Add(1)
		go r.processStream(streamName, groupName, entry)
	}

}

func (r *redisObject) processStream(streamName string, consumerGroup string, stream redis.XMessage) error {
	fmt.Println(stream.Values)
	return r.ro.XAck(streamName, consumerGroup, stream.ID).Err()
}

func (r *redisObject) InitXGroup(streamName string, consumerGroup string) error {
	if _, err := r.ro.XGroupCreateMkStream(streamName, consumerGroup, "0").Result(); err != nil {

		if strings.Contains(fmt.Sprint(err), "BUSYGROUP") {
			// fmt.Printf("Error on create Consumer Group: %v ...\n", consumerGroup)
			// panic(err)
			return nil
		}

		return err
	}

	return nil
}
