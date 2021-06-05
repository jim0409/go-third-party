# intro
使用maxminddb查詢指定ip的地理位置

# intallation
> go get github.com/oschwald/maxminddb-golang

# accerlate query maxminddb speed
```
➜  maxminddb git:(master) ✗ go test . -bench . -cpuprofile prof.cpu 
goos: darwin
goarch: amd64
pkg: go-third-party/maxminddb
BenchmarkLookupSingleIp-8                          47965             28147 ns/op
BenchmarkNetworkLookupSingleIp-8                   44449             30449 ns/op
BenchmarkMapLookupSingleIp-8                    10974222               108 ns/op
BenchmarkLRULookupSingleIp-8                     7813798               152 ns/op
BenchmarkMapLookupWithin_1000_Ips-8              8377029               124 ns/op
BenchmarkLRULookupWithin_1000_Ips-8              7047337               152 ns/op
BenchmarkMapLookupWithin_10000_Ips-8             8302366               133 ns/op
BenchmarkLRULookupWithin_10000_Ips-8             7197708               158 ns/op
BenchmarkMapLookupWithin_100000_Ips-8            9127126               129 ns/op
BenchmarkLRULookupWithin_100000_Ips-8            7834485               152 ns/op
PASS
ok      go-third-party/maxminddb 20.402s
```

<!-- # TODO:
```
happy to know that ... cache map do really accerlate the speed of query

however,
1. as increment of ips in map, down with the query speed
> need to redefine the data

2. over map size would cause broken ...
> need to refactor the container data schema

: (
``` -->
# Fixed
Due to benchmark would have a pre-operation time for each testing. Need to declare huge ip list before use it!

# refer
- maxminddb
> https://github.com/oschwald/maxminddb-golang
- generate random ip
> https://blog.csdn.net/qq_39968176/article/details/90405521