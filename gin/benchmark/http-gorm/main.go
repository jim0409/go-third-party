package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
Running 5s test @ http://127.0.0.1:8000/benchmark/create
  10 threads and 10 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    10.36ms    5.86ms  67.64ms   93.92%
    Req/Sec   102.29     19.52   191.00     75.15%
  5133 requests in 5.10s, 696.76KB read
Requests/sec:   1005.57
Transfer/sec:    136.50KB
*/

/*
Running 5s test @ http://127.0.0.1:8000/benchmark/update
  10 threads and 10 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    25.35ms    5.68ms  82.58ms   92.22%
    Req/Sec    39.64      6.21    50.00     65.60%
  1985 requests in 5.01s, 269.45KB read
Requests/sec:    395.98
Transfer/sec:     53.75KB
*/

var opdb OPDB

func Init() {
	var db = NewDBConfiguration("root", "root", "mysql", "benchmark", "3306", "127.0.0.1")
	idb, err := db.NewDBConnection()
	if err != nil {
		panic(err)
	}

	opdb = idb
}

func main() {
	Init()

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
		v1.GET("create", benchmarkCreateHandler)
		v1.GET("update", benchmarkUpdateHandler)

	}
}

var count = 0

func benchmarkCreateHandler(c *gin.Context) {
	if err := opdb.Create("jim", 1); err != nil {
		c.JSON(500, gin.H{
			"meessage": fmt.Sprintf("%v", err),
		})
		return
	}
	count = count + 1
	c.JSON(200, gin.H{
		"message": "ok",
	})
}

func benchmarkUpdateHandler(c *gin.Context) {
	if err := opdb.Update("jim", count); err != nil {
		c.JSON(500, gin.H{
			"meessage": fmt.Sprintf("%v", err),
		})
		return
	}
	count = count + 1
	c.JSON(200, gin.H{
		"message": "ok",
	})
}
