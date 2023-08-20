package main

import (
	"sync"
	"testing"
)

var (
	pkgimap    = new(IntMap)
	builtinmap = new(sync.Map)
)

// go test -benchmem -run=. -bench .

func BenchmarkPkgSyncMapStore(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pkgimap.Store(i, i)
	}
}

func BenchmarkPkgSyncMapLoad(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pkgimap.Load(i)
	}
}

func BenchmarkBuiltinSyncMapStore(b *testing.B) {
	for i := 0; i < b.N; i++ {
		builtinmap.Store(i, i)
	}
}

func BenchmarkBuiltinSyncMapLoad(b *testing.B) {
	for i := 0; i < b.N; i++ {
		builtinmap.Load(i)
	}
}
