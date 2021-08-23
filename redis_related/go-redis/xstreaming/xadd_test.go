package main

import (
	"strings"
	"testing"

	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
)

var brocaststream = "boradcaststream"

func NewRedisInstance() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "yourpassword",
		DB:       0,
	})
}

func TestGroupMemberPush(t *testing.T) {
	memberIds := []string{}
	memberId := "usc4d2ait94814v8afap9g"
	for i := 0; i < 2000; i++ {
		memberIds = append(memberIds, memberId)
	}
	err := MemberPush("123", strings.Join(memberIds, ","))
	assert.Nil(t, err)
}

func MemberPush(groupID string, members string) error {
	value := map[string]interface{}{
		"Title":    "testtitle",
		"Body":     "testbody",
		"Type":     "testType",
		"ThreadID": "testId",
		"Badge":    "testBadge",
		"Members":  members,
	}

	rds := NewRedisInstance()

	return rds.XAdd(&redis.XAddArgs{
		Stream: brocaststream,
		Values: value,
	}).Err()
}
