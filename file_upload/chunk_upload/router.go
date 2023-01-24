package main

import (
	"github.com/gin-gonic/gin"
)

func apiRouter(router *gin.Engine) {

	route := router.Group("/")
	route.GET("_health", func(c *gin.Context) { c.JSON(200, gin.H{"msg": "hello"}) })

	files := route.Group("/file")
	{
		files.GET("/upload", CheckUploadFile)
		files.POST("/upload", UploadFile)
	}

}

