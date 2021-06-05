package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"go-third-party/gin/reverse_proxy_test/router"
)

func main() {
	//catch global panic
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic err: ", err)
		}
	}()

	route := gin.Default()
	router.ApiRouter(route)

	httpSrv := &http.Server{
		// Addr:    "127.0.0.1:8081",
		Addr:    "0.0.0.0:8081",
		Handler: route,
	}

	go func() {
		if err := httpSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Print(fmt.Sprintf("http listen : %v\n", err))
			panic(err)
		}
	}()

	gracefulShutdown()

}

// gracefulShutdown: handle the worker connection
func gracefulShutdown() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Print()
		fmt.Print(sig)
		// StopDispatcher()
		done <- true
	}()

	// Log.Info("awaiting signal")
	log.Print("awaiting signal")
	<-done
	log.Print("exiting")
}
