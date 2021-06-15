package main

import (
	"testing"

	"github.com/casbin/casbin/v2"
	"gotest.tools/assert"
)

var e = &casbin.Enforcer{}

func init() {
	e, _ = casbin.NewEnforcer("./model.conf", "./policy.csv")
}
func TestCheck(t *testing.T) {
	ok, _ := check(e, "root", "data1", "read")
	assert.Equal(t, true, ok)
}

/*
	BenchmarkCheck-16
	78309 in 1.653s
	15180 ns/op
	7481 B/op
	135 allocs/op
*/
func BenchmarkCheck(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		check(e, "root", "data1", "read")
	}
}
