package main

import (
	"net"
	"sync"

	"github.com/golang/groupcache/lru"
	"github.com/oschwald/maxminddb-golang"
)

type lruCache struct {
	lock     sync.Mutex
	LRUCache *lru.Cache
	Look     *maxminddb.Reader
}

type LRUImpl interface {
	Lookup(string) (map[string]interface{}, error)
}

func WrapLookup(l *maxminddb.Reader) LRUImpl {
	return &lruCache{
		LRUCache: lru.New(1000),
		Look:     l,
	}
}

// Rewrite Lookup with LRU cache ... still failed
// func (l *lruCache) Lookup(ip net.IP, i interface{}) error {
func (l *lruCache) Lookup(ip string) (map[string]interface{}, error) {
	if v, ok := l.LRUCache.Get(ip); ok {
		// log.Println("hit cache")
		return v.(map[string]interface{}), nil
	}

	mapVal := make(map[string]interface{}, 0)
	// log.Println("miss cache")
	var r City
	if err := l.Look.Lookup(net.ParseIP(ip), &r); err != nil {
		return nil, err
	}
	mapVal["city"] = r.City.Names["en"]
	mapVal["country"] = r.Country.Names["en"]
	mapVal["iso"] = r.Country.IsoCode
	mapVal["lat"] = r.Location.Latitude
	mapVal["lgt"] = r.Location.Longitude

	l.lock.Lock()
	l.LRUCache.Add(ip, mapVal)
	l.lock.Unlock()

	return mapVal, nil
}

type City struct {
	City struct {
		GeoNameID uint              `maxminddb:"geoname_id"`
		Names     map[string]string `maxminddb:"names"`
	} `maxminddb:"city"`
	Continent struct {
		Code      string            `maxminddb:"code"`
		GeoNameID uint              `maxminddb:"geoname_id"`
		Names     map[string]string `maxminddb:"names"`
	} `maxminddb:"continent"`
	Country struct {
		GeoNameID         uint              `maxminddb:"geoname_id"`
		IsInEuropeanUnion bool              `maxminddb:"is_in_european_union"`
		IsoCode           string            `maxminddb:"iso_code"`
		Names             map[string]string `maxminddb:"names"`
	} `maxminddb:"country"`
	Location struct {
		AccuracyRadius uint16  `maxminddb:"accuracy_radius"`
		Latitude       float64 `maxminddb:"latitude"`
		Longitude      float64 `maxminddb:"longitude"`
		MetroCode      uint    `maxminddb:"metro_code"`
		TimeZone       string  `maxminddb:"time_zone"`
	} `maxminddb:"location"`
	Postal struct {
		Code string `maxminddb:"code"`
	} `maxminddb:"postal"`
	RegisteredCountry struct {
		GeoNameID         uint              `maxminddb:"geoname_id"`
		IsInEuropeanUnion bool              `maxminddb:"is_in_european_union"`
		IsoCode           string            `maxminddb:"iso_code"`
		Names             map[string]string `maxminddb:"names"`
	} `maxminddb:"registered_country"`
	RepresentedCountry struct {
		GeoNameID         uint              `maxminddb:"geoname_id"`
		IsInEuropeanUnion bool              `maxminddb:"is_in_european_union"`
		IsoCode           string            `maxminddb:"iso_code"`
		Names             map[string]string `maxminddb:"names"`
		Type              string            `maxminddb:"type"`
	} `maxminddb:"represented_country"`
	Subdivisions []struct {
		GeoNameID uint              `maxminddb:"geoname_id"`
		IsoCode   string            `maxminddb:"iso_code"`
		Names     map[string]string `maxminddb:"names"`
	} `maxminddb:"subdivisions"`
	Traits struct {
		IsAnonymousProxy    bool `maxminddb:"is_anonymous_proxy"`
		IsSatelliteProvider bool `maxminddb:"is_satellite_provider"`
	} `maxminddb:"traits"`
}
