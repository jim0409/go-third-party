package main

import (
	"testing"

	"github.com/rs/xid"
)

func BenchmarkRandKey(b *testing.B) {
	for i := 0; i < b.N; i++ {
		xid.New().String()
	}
}
