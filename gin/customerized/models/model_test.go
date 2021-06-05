package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
	redisdb "go-third-party/gin/customerized/models/redisdb"
)

func TestRedisDBAccessImple(t *testing.T) {
	// 讀取並且啟動 redis pool
	redisdb.LoadRedisDBConfig("127.0.0.1:6379", "yourpassword", 300, 500, 1000)
	assert.NoError(t, redisdb.StartRedisPool())

	r := RetriveRedisAccessModel()
	assert.NoError(t, r.SetKey("jim", 123))
	assert.NoError(t, r.DeleteKey("jim"))
}

func TestMySQLDBAccessImple(t *testing.T) {
	ins1 := RetriveRedisAccessModel()
	ins2 := RetriveRedisAccessModel()
	assert.Equal(t, ins1, ins2)
}
