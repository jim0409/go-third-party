package freeport

import (
	"fmt"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFreePort(t *testing.T) {
	port, err := GetFreePort()
	assert.NoError(t, err)
	assert.NotEqual(t, 0, port)
	tmpStr := fmt.Sprintf("localhost:%d", port)
	l, err := net.Listen("tcp", tmpStr)
	assert.NoError(t, err)
	defer l.Close()
}
