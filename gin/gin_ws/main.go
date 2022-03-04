package main

import (
	"log"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("panic err: %v___%v\n", err, string(debug.Stack()))
		}
	}()

	route := gin.Default()
	AddRoute(route)

	httpSrv := &http.Server{
		Addr:    ":8000",
		Handler: route,
	}
	go func() {
		if err := httpSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("http listen err: %v\n", err)
			panic(err)
		}
	}()

	select {}
}

func AddRoute(r *gin.Engine) {
	ws := r.Group("/ws")
	{
		ws.GET("", WebsocketHandler)
	}

}
