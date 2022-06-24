package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

/*
Running 5s test @ http://127.0.0.1:8000/benchmark
  10 threads and 10 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    76.86ms   18.06ms 185.37ms   73.46%
    Req/Sec    13.01      4.81    30.00     68.30%
  647 requests in 5.01s, 87.83KB read
Requests/sec:    129.03
Transfer/sec:     17.52KB
*/

var opdb OPDB

func Init() {
	dbconfig := &DBConfig{
		username:              "root",
		password:              "password",
		address:               "127.0.0.1:27017",
		maxConnectionIdleTime: 5,
		maxPoolSize:           200,
	}

	client, err := dbconfig.connectToMongoDB()
	if err != nil {
		panic(err)
	}

	opdb = OPDB{Client: client.Database("mongo")}

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
	r.GET("benchmark", benchmarkCreateHandler)
	// v1 := r.Group("benchmark")
	// {
	// 	v1.GET("create", benchmarkCreateHandler)
	// 	v1.GET("update", benchmarkUpdateHandler)

	// }
}

var count = 0

func benchmarkCreateHandler(c *gin.Context) {
	account := xid.New().String()
	if err := opdb.Create(account, 1); err != nil {
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

// func benchmarkUpdateHandler(c *gin.Context) {
// 	if err := opdb.Update("jim", count); err != nil {
// 		c.JSON(500, gin.H{
// 			"meessage": fmt.Sprintf("%v", err),
// 		})
// 		return
// 	}
// 	count = count + 1
// 	c.JSON(200, gin.H{
// 		"message": "ok",
// 	})
// }
