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
    Latency     1.29ms    2.08ms  41.34ms   98.88%
    Req/Sec     0.87k   143.51     1.08k    85.53%
  40787 requests in 5.10s, 5.41MB read
Requests/sec:   7990.07
Transfer/sec:      1.06MB
*/

var rsc = NewRedisInstance()

func Init() {
	if err := rsc.HSet("benchmark", "score", 0); err != nil {
		panic(err)
	}

	// if err := rsc.Set("benchmark", 0); err != nil {
	// 	panic(err)
	// }

	// if err := rsc.Lpush("benchmark", 1); err != nil {
	// 	panic(err)
	// }
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
		v1.GET("", benchmarkHandler)

	}
}

var counter = 0

func benchmarkHandler(c *gin.Context) {
	// counter = counter + 1
	// filed := fmt.Sprintf("%v", counter+1)

	if err := rsc.HIncr("benchmark", "score", 1); err != nil {
		c.JSON(500, gin.H{
			"message": err,
		})
		return
	}

	// if err := rsc.Incr("benchmark", 1); err != nil {
	// 	c.JSON(500, gin.H{
	// 		"message": err,
	// 	})
	// 	return
	// }

	// if err := rsc.HSet("benchmark", filed, 1); err != nil {
	// 	c.JSON(500, gin.H{
	// 		"message": err,
	// 	})
	// 	return
	// }

	// if err := rsc.Lpush("benchmark", 1); err != nil {
	// 	c.JSON(500, gin.H{
	// 		"message": err,
	// 	})
	// 	return
	// }

	c.JSON(200, gin.H{
		"message": "ok",
	})
}

func sumup(c *gin.Context) {

}
