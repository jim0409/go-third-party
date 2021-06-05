package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueryIp(t *testing.T) {
	ip := "8.8.8.8"

	locator := LoadIpip(datPath)

	assert.NoError(t, locator.NewIpipObj())

	info, err := locator.QueryIP(ip)
	assert.Equal(t, "GOOGLE", info[0])
	assert.Equal(t, "GOOGLE", info[1])
	assert.Equal(t, "N/A", info[2])
	assert.Equal(t, "N/A", info[3])
	assert.NoError(t, err)
}
