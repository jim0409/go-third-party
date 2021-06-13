package main

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type body struct {
	Message string `json:"message"`
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Create ... 應該要替換func(c *gin.Context)為實際上希望Create的檔案
	// curl -XPOST http://127.0.0.1/test -d '{"message": "test"}'
	router.POST("/test", func(c *gin.Context) {
		b, err := c.GetRawData()
		if err != nil {
			c.JSON(400, gin.H{
				"message": "wrong request data while taking",
			})
			c.Abort()
			return
		}

		pb := body{}
		err = json.Unmarshal(b, &pb)
		if err != nil {
			c.JSON(400, gin.H{
				"message": "wrong request data while json unmarshal",
			})
			c.Abort()
			return
		}

		c.JSON(200, gin.H{
			"message": pb.Message,
		})
	})

	// Read ... 一般的GET請求會直接拿到對應的資料
	// curl -XGET http://127.0.0.1/test?get_query=test
	router.GET("/test", func(c *gin.Context) {
		getStr := c.DefaultQuery("get_query", "test_Read_with_GET_method")
		c.JSON(200, gin.H{
			"message": getStr,
		})
	})

	// Update ... 對應到curl裡面的PUT方法
	router.PUT("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test_Update_with_PUT_method",
		})
	})

	// Delete ... 對應到curl裡面的DELETE方法
	router.DELETE("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test_Delete_with_DELETE_method",
		})
	})

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, map[string]int{
			"status": 1234567,
		})
	})

	router.GET("/ws", wsHandler)

	return router
}

func wsHandler(c *gin.Context) {
	// websocket的Upgrade提供一個Conn: 及其方法 (c *Conn) ...
	w := c.Writer
	r := c.Request

	conn, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
	if err != nil {
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
	}
	go pipelineMsg(conn)
}

func pipelineMsg(c *websocket.Conn) {
	_, msg, err := c.ReadMessage()
	if err != nil && err != websocket.ErrCloseSent {
		return
	}

	if err := c.WriteJSON(string(msg)); err != nil {
		return
	}
}

func main() {
	route := setupRouter()

	httpSrv := &http.Server{
		Addr:    ":80",
		Handler: route,
	}

	httpSrv.ListenAndServe()
}
