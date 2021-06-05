package main

import (
	"testing"

	"github.com/alicebob/miniredis"
	"github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {
	mr, err := miniredis.Run()
	assert.Nil(t, err)

	ro := NewRedisClient(mr.Addr(), "")
	assert.Equal(t, "PONG", ro.ping())

	err = ro.setValue("a", "jim")
	assert.Nil(t, err)

	qstr := ro.queryValue("a")
	assert.Equal(t, "Query key 'a', get return value 'jim'\n", qstr)

	err = ro.deleteValue("a")
	assert.Nil(t, err)

	qstr = ro.queryValue("a")
	assert.Equal(t, "Error happend with redis: nil\n", qstr)
}
