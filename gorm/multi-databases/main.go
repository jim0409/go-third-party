package main

import (
	"fmt"
	"go-third-party/gorm/multi-databases/models"
	"math/rand"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	usr    = os.Args[1] // "jim"
	pwd    = os.Args[2] // "password"
	dbt    = os.Args[3] //"mysql"
	dbname = os.Args[4] //"db"
	port   = os.Args[5] // "3306"
	addr   = os.Args[6] //"127.0.0.1"
	// go run main.go jim password mysql db 3306 127.0.0.1
)

// 應該跟 DB 決策器綁定一起，透過決策器回饋的數字。產生出對應的值
var rGen = func() (int, string, string, string) {
	id := rand.Intn(3)
	group := fmt.Sprintf("msg-%d", id)
	name := fmt.Sprintf("jim_%d", id)
	age := fmt.Sprintf("3%d", id)
	return id, group, name, age
}

func main() {
	m := models.InitMainDB(usr, pwd, dbt, dbname, port, addr)
	defer func() {
		if err := m.Closed(); err != nil {
			panic(err)
		}
	}()

	router := gin.Default()
	router.GET("/purge", func(c *gin.Context) {
		if err := m.PurgeCache(); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusOK, "purge cache")
	})

	router.GET("/insert", func(c *gin.Context) {
		i, g, n, a := rGen()
		if err := m.CreateMessage(g, n, a); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"id":     fmt.Sprintf("%d", i),
			"status": "ok",
		})
	})

	router.GET("/read", func(c *gin.Context) {
		_, g, _, _ := rGen()
		res, err := m.ReadMessage(g, nil)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"result": res,
		})
	})

	httpSrv := &http.Server{
		Addr:    ":8001",
		Handler: router,
	}

	if err := httpSrv.ListenAndServe(); err != nil {
		panic(err)
	}
}
