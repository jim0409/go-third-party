package main

import "testing"

func BenchmarkHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		genHash("a4b271d2fd9022110eeca555e9011f2863ae36b8473ad70797528c2de6491aea")
	}
}
