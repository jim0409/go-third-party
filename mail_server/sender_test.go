package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSender(t *testing.T) {
	cfg, err := InitConfig("./config.ini")
	if err != nil {
		panic(err)
	}
	SetGlobalVariable(*cfg)
	sender := NewSender(GlobalAuth, GlobalUser, GlobalHost, GlobalSmtpserver, &MsgTemplate)

	err = sender.SendMail(DemoFrom, DemoTo, DemoSubject, DemoId)
	assert.Nil(t, err)
}
