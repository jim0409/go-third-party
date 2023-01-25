package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var opdb OPDB

func main() {
	// check mysql connection
	mysqlAddr := "127.0.0.1"
	mysqlPort := "3306"
	mysqlOpDB := "upload_file"
	mysqlUsr := "root"
	mysqUsrPwd := "root"

	newDB := NewDBConfiguration(mysqlUsr, mysqUsrPwd, "mysql", mysqlOpDB, mysqlPort, mysqlAddr)
	db, err := newDB.NewDBConnection()
	if err != nil {
		log.Fatal(err)
	}
	opdb = db

	// check file save folder
	if _, err := os.Stat("files"); os.IsNotExist(err) {
		if err := os.Mkdir("files", 0777); err != nil {
			// return 0, err
			log.Fatal(err)
		}
	}

	router := gin.Default()
	apiRouter(router)

	httpSrv := &http.Server{
		Addr:    ":" + "8000",
		Handler: router,
	}

	go func() {
		if err := httpSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("http listen : %v\n", err)
			panic(err)
		}
	}()

	fmt.Printf("==== Now server working on %v ====\n", httpSrv.Addr)

	select {}
}
