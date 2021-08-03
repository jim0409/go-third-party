package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	mysqlAddr := "127.0.0.1"
	mysqlPort := "3306"
	mysqlOpDB := "testdb"
	mysqlUsr := "jim"
	mysqUsrPwd := "password"

	newDB := NewDBConfiguration(mysqlUsr, mysqUsrPwd, "mysql", mysqlOpDB, mysqlPort, mysqlAddr)
	db, err := newDB.NewDBConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := db.Closed(); err != nil {
			log.Fatal(err)
		}
	}()

	router := gin.Default()
	router.GET("/insert", func(c *gin.Context) {
		db.create("jim", "mail")
		c.JSON(200, "ok")
	})

	httpSrv := &http.Server{
		Addr:    ":80",
		Handler: router,
	}

	httpSrv.ListenAndServe()

}
