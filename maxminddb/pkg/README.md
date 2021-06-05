# intro
build a packet for maxminddb dao with default LRU cached

# GOAL
> retrive city info via ip

# benchmark cli
> go test . -bench . -cpuprofile prof.cpu 
```log
➜  pkg git:(master) ✗ go test . -bench . -cpuprofile prof.cpu 
goos: darwin
goarch: amd64
pkg: go-third-party/maxminddb/pkg
BenchmarkLRUGeoCache-8                  26733382                42.1 ns/op
BenchmarkMapLookupWithin_1000_Ips-8     23124248                47.9 ns/op
PASS
ok      go-third-party/maxminddb/pkg     5.982s
```

