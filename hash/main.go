package main

import (
	"fmt"
	"hash/maphash"
	"log"
)

var (
	hash     maphash.Hash
	hashSeed = hash.Seed()
)

func genHash(id string) uint64 {
	b := []byte(id)
	hash.SetSeed(hashSeed)
	_, err := hash.Write(b)
	if err != nil {
		log.Fatal(err)
	}
	return hash.Sum64() % uint64(5)

}

func main() {
	userID := "a4b271d2fd9022110eeca555e9011f2863ae36b8473ad70797528c2de6491aea"
	u64 := genHash(userID)
	fmt.Println(u64)
}
