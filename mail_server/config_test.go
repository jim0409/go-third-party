package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadConfig(t *testing.T) {
	cfg, err := InitConfig("./config.ini")
	assert.Nil(t, err)
	fmt.Println(cfg.BaseConf)
	fmt.Println(cfg.SmtpConf)
	fmt.Println(cfg.User)
}
