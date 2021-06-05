package main

import (
	"net"
	"sync"

	"github.com/golang/groupcache/lru"
)

type lruCacheLookUp struct {
	lock     sync.Mutex
	LRUCache *lru.Cache

	Look LRULookup
}

type LRULookup interface {
	Lookup(net.IP, interface{}) error
}

func LRUWrapLookup(l LRULookup) LRULookup {
	return &lruCacheLookUp{
		LRUCache: lru.New(1000),
		Look:     l,
	}
}

// Rewrite Lookup with LRU cache ... still failed
func (l *lruCacheLookUp) Lookup(ip net.IP, i interface{}) error {
	var queryIp = ip.String()
	if v, ok := l.LRUCache.Get(queryIp); ok {
		// log.Println("hit cache")
		i = v
		return nil
	} else {
		// log.Println("miss cache")
		if err := l.Look.Lookup(ip, &i); err != nil {
			return err
		}
		l.lock.Lock()
		l.LRUCache.Add(queryIp, i)
		l.lock.Unlock()
	}

	return nil
}
