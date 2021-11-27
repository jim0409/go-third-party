package main

import (
	"fmt"
	"net/http"
)

/*
➜  go-third-party git:(main) ✗ wrk -t100 -c100 http://127.0.0.1:3002/
Running 10s test @ http://127.0.0.1:3002/
  100 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.09ms  348.99us   5.22ms   79.77%
    Req/Sec     0.92k    81.98     1.49k    82.49%
  929609 requests in 10.10s, 120.57MB read
Requests/sec:  92005.25
Transfer/sec:     11.93MB
*/

func muxServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, World 👋!", "\n")

	})
	go http.ListenAndServe(":3002", mux)
}
