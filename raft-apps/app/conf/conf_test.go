package conf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConf(t *testing.T) {
	path := "./app.dev.ini"
	config, err := InitConfig(path)
	assert.Nil(t, err)
	assert.Equal(t, 16251, config.HttpPort)
	assert.Equal(t, "http://127.0.0.1:12379", config.PeerAddr)
	assert.Equal(t, "dev", config.Env)
	assert.Equal(t, "raft", config.DbName)
	assert.Equal(t, "127.0.0.1", config.DbHost)
	assert.Equal(t, "3306", config.DbPort)
	assert.Equal(t, "raft", config.DbUser)
	assert.Equal(t, "raft", config.DbPassword)
	assert.Equal(t, false, config.DbLogEnable)
	assert.Equal(t, 300, config.DbMaxConnect)
	assert.Equal(t, 10, config.DbIdleConnect)
}
