package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLrange(t *testing.T) {
	rds := NewRedisClient("127.0.0.1:6379", "yourpassword")
	strs, err := rds.lrange("cluster")
	assert.Nil(t, err)
	fmt.Println(strs)
}
