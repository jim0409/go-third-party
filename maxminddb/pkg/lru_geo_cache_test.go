package main

import (
	"fmt"
	"log"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var InitIPSet []string

func init() {
	InitIPSet = setRandIpArr(100000)
	// fmt.Println(len(InitIPSet))
}

func setRandIpArr(n int) []string {
	var ips = make([]string, 0)

	// generate rand ips
	for i := 0; i < n; i++ {
		ips = append(ips, genIpaddr())
		// log.Printf("%d_%v\t", i, ips[i]) // verified that ip differ from each other
	}
	return ips
}

func genIpaddr() string {
	rand.Seed(time.Now().UnixNano())
	ip := fmt.Sprintf("%d.%d.%d.%d", rand.Intn(255), rand.Intn(255), rand.Intn(255), rand.Intn(255))
	return ip
}

func TestLRUGeoCache(t *testing.T) {
	db, err := connect("../GeoIP2-City-Test.mmdb")
	if err != nil {
		log.Fatal()
	}
	defer db.Close()
	lruDb := WrapLookup(db)

	ip := "81.2.69.142"

	m, err := lruDb.Lookup(ip)
	if err != nil {
		log.Fatal(err)
	}
	// log.Println(m)
	assert.NotNil(t, m)

	m2, err := lruDb.Lookup(ip)
	if err != nil {
		log.Fatal(err)
	}
	// log.Println(m2)
	assert.NotNil(t, m2)
}

func BenchmarkLRUGeoCache(b *testing.B) {
	db, err := connect("../GeoIP2-City-Test.mmdb")
	if err != nil {
		log.Fatal()
	}
	defer db.Close()
	lruDb := WrapLookup(db)

	ip := "81.2.69.142"
	for i := 0; i < b.N; i++ {
		if m, err := lruDb.Lookup(ip); err != nil || m == nil {
			log.Fatal(err)
		}
	}
}

func BenchmarkMapLookupWithin_1000_Ips(b *testing.B) {
	db, err := connect("../GeoIP2-City-Test.mmdb")
	if err != nil {
		log.Fatal()
	}
	defer db.Close()
	lruDb := WrapLookup(db)

	ipList := InitIPSet
	// inject ip to ipList
	for _, ip := range ipList {
		if m, err := lruDb.Lookup(ip); err != nil || m == nil {
			log.Fatal(err)
		}
	}

	ip := "81.2.69.142"
	for i := 0; i < b.N; i++ {
		if m, err := lruDb.Lookup(ip); err != nil || m == nil {
			log.Fatal(err)
		}
	}
}
