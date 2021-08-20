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
	XRead(string)
	XAdd(string, map[string]interface{}) (string, error)
	XReadGroup(string, string, string)
	ConsumePendingGroup(string, string, string)
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

func (r *redisObject) XAdd(streamName string, value map[string]interface{}) (string, error) {
	newId, err := r.ro.XAdd(&redis.XAddArgs{
		Stream: streamName,
		Values: value,
	}).Result()
	if err != nil {
		return "", err
	}

	return newId, nil
}

func (r *redisObject) XRead(streamName string) {
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

	// del id would cause a strictly performance cost
	for _, msg := range entries[0].Messages {
		fmt.Println(msg)
		fmt.Println(r.ro.XDel(streamName, msg.ID).Result())
	}
}

// xread 與 xgroupread 差異:
// http://www.redis.cn/commands/xreadgroup.html
func (r *redisObject) XReadGroup(streamName string, groupName string, consumerName string) {
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

	// _ = entries
	for _, entry := range entries[0].Messages {
		// waitGrp.Add(1)
		go r.processStream(streamName, groupName, entry)
	}

}

func (r *redisObject) processStream(streamName string, consumerGroup string, stream redis.XMessage) error {
	fmt.Println(r.len(streamName), stream.ID, stream.Values)
	return r.ro.XAck(streamName, consumerGroup, stream.ID).Err()
}

// 消費一些逾期尚未處理的 streaming message
func (r *redisObject) ConsumePendingGroup(streamName string, consumerGroup string, consumerName string) {
	var streamsRetry []string
	pendingStreams, err := r.ro.XPendingExt(&redis.XPendingExtArgs{
		Stream: streamName,
		Group:  consumerGroup,
		Start:  "0",
		End:    "+",
		Count:  10,
		//Consumer string
	}).Result()

	if err != nil {
		panic(err)
	}

	for _, stream := range pendingStreams {
		streamsRetry = append(streamsRetry, stream.Id)
	}

	if len(streamsRetry) > 0 {

		streams, err := r.ro.XClaim(&redis.XClaimArgs{
			Stream:   streamName,
			Group:    consumerGroup,
			Consumer: consumerName,
			Messages: streamsRetry,
			MinIdle:  30 * time.Second,
		}).Result()

		if err != nil {
			log.Printf("err on process pending: %+v\n", err)
			return
		}

		for _, stream := range streams {
			// waitGrp.Add(1)
			go r.processStream(streamName, consumerGroup, stream)
		}
		// waitGrp.Wait()
	}

	fmt.Println("process pending streams at ", time.Now().Format(time.RFC3339))
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
