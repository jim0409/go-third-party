package main

import (
	"bytes"
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

func main() {

	db, err := connect("../GeoIP2-City-Test.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ip := net.ParseIP("81.2.69.142")

	var record City
	db.Lookup(ip, &record)

	trimF := func(a string) string {
		b := bytes.Buffer{}
		for i := 0; i < len(a); i++ {
			if a[i:i+1] != " " {
				b.WriteString(a[i : i+1])
			}
		}
		return b.String()
	}

	log.Printf("%v_%v_%v_%v_%v\n",
		record.City.Names["en"],
		trimF(record.Country.Names["en"]),
		record.Country.IsoCode,
		record.Location.Latitude,
		record.Location.Longitude,
	)

}
