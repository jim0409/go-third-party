package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

/*
go-redis在對cluster的情況下，因為每次訪問的節點都不同，
所以設定鍵值前要先找該節點，也因此用同樣的方法請求都會偏慢...
要用pipeline才能加速cluster的運作...
*/

var rsc = NewRedisInstance()

func Init() {

	if err := rsc.Set("benchmark", 1); err != nil {
		panic(err)
	}
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

	key := xid.New().String()
	if err := rsc.Set(key, 1); err != nil {
		c.JSON(500, gin.H{
			"message": err,
		})
		return
	}

	defer rsc.Client.Pipeline().Exec()

	c.JSON(200, gin.H{
		"message": "ok",
	})
}
