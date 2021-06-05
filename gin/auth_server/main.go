package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	apiRouter(router)

	// router.Run() // usually used for dev env

	httpSrv := &http.Server{
		Addr:    ":" + "9000",
		Handler: router,
	}

	go func() {
		if err := httpSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("http listen : %v\n", err)
			panic(err)
		}
	}()

	fmt.Printf("==== Now server working on %v ====\n", httpSrv.Addr)

	// deal with graceful shutdown
	gracefulShutdown()

}

func apiRouter(router *gin.Engine) {

	authorized := router.Group("/")

	authorized.Use(AuthCheck(), AddLog())

	ver1 := authorized.Group("/test")
	{
		ver1.GET("g1", helloHandler)
		ver1.GET("g2", CheckParamAndHeader(helloHandler))
		// ver1.POST("g3", AddLog(helloHandler))
		ver1.POST("g3", helloHandler)
	}

}

func helloHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello",
	})
}

func gracefulShutdown() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}

// CheckParamAndHeader是一個裝飾器方法，透過傳入的`Context`回傳對應的Handler
func CheckParamAndHeader(h gin.HandlerFunc) gin.HandlerFunc {
	f := func(c *gin.Context) {
		header := c.Request.Header.Get("token")
		if header == "" {
			c.JSON(400, gin.H{
				"code":   "3",
				"result": "failed",
				"msg":    "Missing token",
			})
			return
		} else {
			h(c)
		}
	}

	return f

}

// UserMap為假定，合法的使用者
var UserMap = map[string]int{"jim": 1, "tim": 1}

// AuthCheck是一個前置檢查動作，所有請求都需要滿足的
func AuthCheck() gin.HandlerFunc {
	f := func(c *gin.Context) {
		user := c.GetHeader("User")
		switch UserMap[user] {
		case 1:
			return
		default:
			invalid(c)

		}
	}
	return f
}

// 非法使用者的handlerFunc
func invalid(c *gin.Context) {
	c.JSON(400, gin.H{
		"code":   "2",
		"result": "failed",
		"msg":    "Invalid User.",
	})
	c.Abort()
	return
}

// func AddLog(h gin.HandlerFunc) gin.HandlerFunc {
// 	f := func(c *gin.Context) {
// 		buf := make([]byte, 1024)
// 		num, _ := c.Request.Body.Read(buf)
// 		reqBody := string(buf[0:num])
// 		// Log.Info("Receiv Raw Request ", reqBody)
// 		log.Println(reqBody)
// 		h(c)
// 	}
// 	return f
// }

func AddLog() gin.HandlerFunc {
	f := func(c *gin.Context) {
		var bodyBytes []byte
		if c.Request.Context() != nil {
			bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
		}
		log.Printf("Receive Raw Request body with %v\n", string(bodyBytes))

		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	}
	return f
}
