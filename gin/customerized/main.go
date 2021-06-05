package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"go-third-party/gin/customerized/conf"
	"go-third-party/gin/customerized/router"

	. "go-third-party/gin/customerized/internal"
	. "go-third-party/gin/customerized/logger"
	. "go-third-party/gin/customerized/models"
)

var (
	HTTPAddr    = flag.String("http", "127.0.0.1:9090", "Address to listen for HTTP requests on")
	confPath    = flag.String("config", "./conf/app.dev.ini", "config location")
	checkcommit = flag.Bool("version", false, "burry code for check version")

	confInfo     *conf.Config
	gitcommitnum string
)

func checkComimit() {
	log.Println(gitcommitnum)
}

func Init() error {
	flag.Parse()
	// if there is a needed to check git commit num ... print it out
	if *checkcommit {
		checkComimit()
		os.Exit(1)
	}

	// read config and pass variables ...
	var err error
	confInfo, err = conf.InitConfig(confPath)
	if err != nil {
		return fmt.Errorf("Init config err: %v", err)
	}

	fmt.Println(confInfo.WorkersPool.Tag, confInfo.WorkersPool.NumWorkers)

	// initialize logger
	if err = InitLog(confInfo.LogConf.LogPath, confInfo.LogConf.LogLevel); err != nil {
		return fmt.Errorf("init log is err: %v", err)
	}

	// initialize both mysql and redis db
	if err = InitDb(&confInfo.DbConf, &confInfo.RedisConf); err != nil {
		return fmt.Errorf("init db is err: %v", err)
	}

	// initialize workerpools
	InitWorkerPools(confInfo.WorkersPool.Tag, confInfo.WorkersPool.NumWorkers)
	return err
}

func main() {
	//catch global panic
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic err: ", err)
		}
	}()

	err := Init()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	route := gin.Default()
	router.ApiRouter(route)

	httpSrv := &http.Server{
		Addr:    *HTTPAddr,
		Handler: route,
	}

	go func() {
		if err := httpSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			Log.Info(fmt.Sprintf("http listen : %v\n", err))
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
		fmt.Println()
		fmt.Println(sig)
		StopDispatcher()
		done <- true
	}()

	Log.Info("awaiting signal")
	<-done
	Log.Info("exiting")
}
