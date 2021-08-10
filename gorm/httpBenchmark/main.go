package main

import (
	"log"
	"net/http"

	dbpkg "go-third-party/gorm"

	"github.com/gin-gonic/gin"
)

func main() {

	mysqlAddr := "127.0.0.1"
	mysqlPort := "3306"
	mysqlOpDB := "mysql"
	mysqlUsr := "root"
	// mysqUsrPwd := "1qaz!QAZ"
	mysqUsrPwd := "secret"

	newDB := dbpkg.NewDBConfiguration(mysqlUsr, mysqUsrPwd, "mysql", mysqlOpDB, mysqlPort, mysqlAddr)
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
		db.Create("jim", "mail")
		c.JSON(http.StatusOK, "ok")
	})

	httpSrv := &http.Server{
		Addr:    ":8001",
		Handler: router,
	}

	if err := httpSrv.ListenAndServe(); err != nil {
		panic(err)
	}

}
