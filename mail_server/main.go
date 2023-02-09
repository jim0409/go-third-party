package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	GlobalAuth       string
	GlobalUser       string
	GlobalHost       string
	GlobalSmtpserver string
	GlobalEnv        string
)

var (
	DemoFrom    = "system@gmail.com"
	DemoTo      = "berserker.01.tw@gmail.com"
	DemoSubject = "system subject"
	DemoId      = 0
)

func SetGlobalVariable(cfg Config) {
	GlobalAuth = cfg.Auth
	GlobalUser = cfg.User
	GlobalHost = cfg.Host
	GlobalSmtpserver = cfg.Server
	GlobalEnv = cfg.Env
}

func main() {
	cfg, err := InitConfig("./config.ini")
	if err != nil {
		panic(err)
	}
	SetGlobalVariable(*cfg)

	handler := gin.Default()
	apiRouter(handler)

	httpSrv := &http.Server{
		Addr:    ":" + "8000",
		Handler: handler,
	}

	go func() {
		if err := httpSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("http listen : %v\n", err)
			panic(err)
		}
	}()

	select {}
}
