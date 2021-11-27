package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
➜  go-third-party git:(main) ✗ wrk -t100 -c100 http://127.0.0.1:3001
Running 10s test @ http://127.0.0.1:3001
  100 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.07ms  387.94us   7.67ms   79.61%
    Req/Sec     0.94k    59.69     1.23k    69.28%
  946268 requests in 10.10s, 121.83MB read
Requests/sec:  93660.60
Transfer/sec:     12.06MB


➜  go-third-party git:(main) ✗ wrk -t100 -c100 http://127.0.0.1:3001/123
Running 10s test @ http://127.0.0.1:3001/123
  100 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.08ms  365.52us   5.21ms   82.45%
    Req/Sec     0.93k    68.89     1.26k    64.72%
  935227 requests in 10.10s, 111.49MB read
Requests/sec:  92556.79
Transfer/sec:     11.03MB
*/

func ginServer() {
	route := gin.New()

	route.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World 👋!")
		// c.JSON(200, gin.H{
		// 	"message": "pong",
		// })
	})

	route.GET("/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.String(200, fmt.Sprintf("%v 👋!", id))
		// c.JSON(200, gin.H{
		// 	"message": "pong",
		// })
	})

	httpSrv := &http.Server{
		Addr:    ":3001",
		Handler: route,
	}
	// r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	go func() {
		if err := httpSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("http listen : %v\n", err)
			panic(err)
		}
	}()

}
