package redisdb

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRedisSetKey(t *testing.T) {
	// 讀取並且啟動 redis pool
	LoadRedisDBConfig("127.0.0.1:6379", "yourpassword", 300, 500, 1000)
	assert.NoError(t, StartRedisPool())

	// 拿取 redis pool的物件
	redispool := RetriveRedisPoolObj()

	// 插入一個 redis key/value
	assert.NoError(t, redispool.SetKey("companyA:jim:456", 1))
	assert.NoError(t, redispool.DeleteKey("companyA:jim:456"))

	// 插入一個 排斥的 redis key
	assert.NoError(t, redispool.SetNonExistedKey("companyA:jim:456", 31))

	assert.Error(t, errors.New("Failed to set non-existed key to redis"), redispool.SetNonExistedKey("companyA:jim:456", 21))

	assert.NoError(t, redispool.SetExpiredKey("companyA:jim:456", 0, 10))
	// assert.NoError(t, redispool.Close())
}
