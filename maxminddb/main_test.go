package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"testing"
	"time"
)

var InitIPSet [](net.IP)

func init() {
	InitIPSet = setRandIpArr(100000)
	// fmt.Println(len(InitIPSet))
}

func setRandIpArr(n int) [](net.IP) {
	var ips = make([](net.IP), 0)

	// generate rand ips
	for i := 0; i < n; i++ {
		ips = append(ips, net.ParseIP(genIpaddr()))
		// log.Printf("%d_%v\t", i, ips[i]) // verified that ip differ from each other
	}
	return ips
}

func genIpaddr() string {
	rand.Seed(time.Now().UnixNano())
	ip := fmt.Sprintf("%d.%d.%d.%d", rand.Intn(255), rand.Intn(255), rand.Intn(255), rand.Intn(255))
	return ip
}

func BenchmarkLookupSingleIp(b *testing.B) {
	db, err := connect("./GeoIP2-City-Test.mmdb")
	if err != nil {
		log.Fatal()
	}
	defer db.Close()
	ip := net.ParseIP("81.2.69.142")
	var record interface{}

	for i := 0; i < b.N; i++ {
		if err := QueryMaxmindDB(db, ip, &record); err != nil {
			log.Fatal(err)
		}
	}
}

func BenchmarkNetworkLookupSingleIp(b *testing.B) {
	db, err := connect("./GeoIP2-City-Test.mmdb")
	if err != nil {
		log.Fatal()
	}
	defer db.Close()
	ip := net.ParseIP("81.2.69.142")
	var record interface{}

	for i := 0; i < b.N; i++ {
		if err := QueryMaxmindDBWithNetwork(db, ip, &record); err != nil {
			log.Fatal(err)
		}
	}

}

func BenchmarkMapLookupSingleIp(b *testing.B) {
	db, err := connect("./GeoIP2-City-Test.mmdb")
	if err != nil {
		log.Fatal()
	}
	defer db.Close()
	ip := net.ParseIP("81.2.69.142")
	var acRecord interface{}
	mapCacaheIpRecord := MapWrapLookup(db)

	for i := 0; i < b.N; i++ {
		if err := mapCacaheIpRecord.Lookup(ip, &acRecord); err != nil {
			log.Fatal(err)
		}
	}
}

func BenchmarkLRULookupSingleIp(b *testing.B) {
	db, err := connect("./GeoIP2-City-Test.mmdb")
	if err != nil {
		log.Fatal()
	}
	defer db.Close()
	ip := net.ParseIP("81.2.69.142")
	var acRecord interface{}
	lruCacaheIpRecord := LRUWrapLookup(db)

	for i := 0; i < b.N; i++ {
		if err := lruCacaheIpRecord.Lookup(ip, &acRecord); err != nil {
			log.Fatal(err)
		}
	}
}

// Simulate that there exist multi-ips in cache map
func BenchmarkMapLookupWithin_1000_Ips(b *testing.B) {
	db, err := connect("./GeoIP2-City-Test.mmdb")
	if err != nil {
		log.Fatal()
	}
	defer db.Close()
	var acRecord interface{}
	mapCacaheIpRecord := MapWrapLookup(db)

	ipList := setRandIpArr(1000)
	// inject ip to ipList
	for _, ip := range ipList {
		if err := mapCacaheIpRecord.Lookup(ip, &acRecord); err != nil {
			log.Fatal(err)
		}
	}

	// test ip
	ip := net.ParseIP("81.2.69.142")

	for i := 0; i < b.N; i++ {
		if err := mapCacaheIpRecord.Lookup(ip, &acRecord); err != nil {
			log.Fatal(err)
		}
	}
}

func BenchmarkLRULookupWithin_1000_Ips(b *testing.B) {
	db, err := connect("./GeoIP2-City-Test.mmdb")
	if err != nil {
		log.Fatal()
	}
	defer db.Close()
	var acRecord interface{}
	lruCacaheIpRecord := LRUWrapLookup(db)

	ipList := setRandIpArr(1000)
	// inject ip to ipList
	for _, ip := range ipList {
		if err := lruCacaheIpRecord.Lookup(ip, &acRecord); err != nil {
			log.Fatal(err)
		}
	}

	// test ip
	ip := net.ParseIP("81.2.69.142")

	for i := 0; i < b.N; i++ {
		if err := lruCacaheIpRecord.Lookup(ip, &acRecord); err != nil {
			log.Fatal(err)
		}
	}
}

func BenchmarkMapLookupWithin_10000_Ips(b *testing.B) {
	db, err := connect("./GeoIP2-City-Test.mmdb")
	if err != nil {
		log.Fatal()
	}
	defer db.Close()
	var acRecord interface{}
	mapCacaheIpRecord := MapWrapLookup(db)

	ipList := setRandIpArr(10000)
	// inject ip to ipList
	for _, ip := range ipList {
		if err := mapCacaheIpRecord.Lookup(ip, &acRecord); err != nil {
			log.Fatal(err)
		}
	}

	// test ip
	ip := net.ParseIP("81.2.69.142")

	for i := 0; i < b.N; i++ {
		if err := mapCacaheIpRecord.Lookup(ip, &acRecord); err != nil {
			log.Fatal(err)
		}
	}
}

func BenchmarkLRULookupWithin_10000_Ips(b *testing.B) {
	db, err := connect("./GeoIP2-City-Test.mmdb")
	if err != nil {
		log.Fatal()
	}
	defer db.Close()
	var acRecord interface{}
	lruCacaheIpRecord := LRUWrapLookup(db)

	ipList := setRandIpArr(10000)
	// inject ip to ipList
	for _, ip := range ipList {
		if err := lruCacaheIpRecord.Lookup(ip, &acRecord); err != nil {
			log.Fatal(err)
		}
	}

	// test ip
	ip := net.ParseIP("81.2.69.142")

	for i := 0; i < b.N; i++ {
		if err := lruCacaheIpRecord.Lookup(ip, &acRecord); err != nil {
			log.Fatal(err)
		}
	}
}

func BenchmarkMapLookupWithin_100000_Ips(b *testing.B) {
	db, err := connect("./GeoIP2-City-Test.mmdb")
	if err != nil {
		log.Fatal()
	}
	defer db.Close()
	var acRecord interface{}
	mapCacaheIpRecord := MapWrapLookup(db)

	// ipList := setRandIpArr(100000)
	ipList := InitIPSet
	// inject ip to ipList
	for _, ip := range ipList {
		if err := mapCacaheIpRecord.Lookup(ip, &acRecord); err != nil {
			log.Fatal(err)
		}
	}

	// test ip
	ip := net.ParseIP("81.2.69.142")

	for i := 0; i < b.N; i++ {
		if err := mapCacaheIpRecord.Lookup(ip, &acRecord); err != nil {
			log.Fatal(err)
		}
	}
}

func BenchmarkLRULookupWithin_100000_Ips(b *testing.B) {
	db, err := connect("./GeoIP2-City-Test.mmdb")
	if err != nil {
		log.Fatal()
	}
	defer db.Close()
	var acRecord interface{}
	lruCacaheIpRecord := LRUWrapLookup(db)

	// ipList := setRandIpArr(100000)
	ipList := InitIPSet
	// inject ip to ipList
	for _, ip := range ipList {
		if err := lruCacaheIpRecord.Lookup(ip, &acRecord); err != nil {
			log.Fatal(err)
		}
	}

	// test ip
	ip := net.ParseIP("81.2.69.142")

	for i := 0; i < b.N; i++ {
		if err := lruCacaheIpRecord.Lookup(ip, &acRecord); err != nil {
			log.Fatal(err)
		}
	}
}

// Consider parallel case
