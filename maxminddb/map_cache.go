package main

import (
	"net"
	"sync"
)

type mapCacheLookUp struct {
	lock     sync.Mutex
	cacheMap map[string]interface{}

	Look MapLookup
}

type MapLookup interface {
	Lookup(net.IP, interface{}) error
}

func MapWrapLookup(l MapLookup) MapLookup {
	return &mapCacheLookUp{
		cacheMap: make(map[string]interface{}, 0),
		Look:     l,
	}
}

func (c *mapCacheLookUp) Lookup(ip net.IP, i interface{}) error {
	var queryIp = ip.String()
	if v, ok := c.cacheMap[queryIp]; ok {
		// log.Println("hit cache")
		i = v
		return nil
	} else {
		// log.Println("miss cache")
		if err := c.Look.Lookup(ip, &i); err != nil {
			return err
		}
		c.lock.Lock()
		c.cacheMap[queryIp] = i
		c.lock.Unlock()
	}

	return nil
}
