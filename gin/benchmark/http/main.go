package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
Running 5s test @ http://127.0.0.1:8000/benchmark
  10 threads and 10 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   454.42us    2.17ms  37.08ms   97.05%
    Req/Sec     9.98k     2.36k   13.59k    67.65%
  506350 requests in 5.10s, 68.57MB read
Requests/sec:  99297.73
Transfer/sec:     13.45MB
*/

func main() {
	router := gin.New()
	apiRouter(router)

	httpSrv := &http.Server{
		Addr:    ":" + "8000",
		Handler: router,
	}

	if err := httpSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Printf("http listen : %v\n", err)
		panic(err)
	}

}

func apiRouter(router *gin.Engine) {
	r := router.Group("/")
	v1 := r.Group("benchmark")
	{
		v1.GET("", benchmarkHandler)

	}
}

func benchmarkHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ok",
	})
}
