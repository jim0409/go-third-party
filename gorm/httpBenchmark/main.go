package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"time"

	dbpkg "go-third-party/gorm"

	"github.com/gin-gonic/gin"
)

func main() {
	net.DefaultResolver = &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{
				Timeout: time.Millisecond * time.Duration(10000),
			}
			return d.DialContext(ctx, network, "127.0.0.1:5301")
		},
	}

	// mysqlAddr := "127.0.0.1"
	mysqlAddr := "dns.jim.host"
	mysqlPort := "3306"
	mysqlOpDB := "mysql"
	mysqlUsr := "root"
	// mysqUsrPwd := "1qaz!QAZ"
	mysqUsrPwd := "root"

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
