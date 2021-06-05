package main

import (
	"log"
	"net"

	"github.com/oschwald/maxminddb-golang"
)

func connect(path string) (*maxminddb.Reader, error) {
	db, err := maxminddb.Open(path)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func QueryMaxmindDB(db *maxminddb.Reader, ip net.IP, i *interface{}) error {
	return db.Lookup(ip, i)
}

func QueryMaxmindDBWithNetwork(db *maxminddb.Reader, ip net.IP, i *interface{}) error {
	_, _, err := db.LookupNetwork(ip, i)
	return err
}

func main() {
	db, err := connect("./GeoIP2-City-Test.mmdb")
	if err != nil {
		log.Fatal()
	}
	defer db.Close()
	ip := net.ParseIP("81.2.69.142")
	var record interface{}

	if err := QueryMaxmindDB(db, ip, &record); err != nil {
		log.Fatal(err)
	}
	log.Println(record)

	var mapRecord interface{}
	mapCacaheIpRecord := MapWrapLookup(db)
	if err := mapCacaheIpRecord.Lookup(ip, &mapRecord); err != nil {
		log.Fatal(err)
	}
	log.Println(mapRecord)

	var lruRecord interface{}
	lruCacaheIpRecord := LRUWrapLookup(db)
	if err := lruCacaheIpRecord.Lookup(ip, &lruRecord); err != nil {
		log.Fatal(err)
	}
	log.Println(lruRecord)

	// check hit/miss for specific case
	ip = net.ParseIP("127.0.0.1")
	var lruPenertationRecord interface{}
	if err := lruCacaheIpRecord.Lookup(ip, &lruPenertationRecord); err != nil {
		log.Fatal(err)
	}
	log.Println(lruRecord)

	if err := lruCacaheIpRecord.Lookup(ip, &lruPenertationRecord); err != nil {
		log.Fatal(err)
	}
	log.Println(lruRecord)
}
