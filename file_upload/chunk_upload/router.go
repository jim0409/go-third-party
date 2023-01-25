package main

import (
	"github.com/gin-gonic/gin"
)

func apiRouter(router *gin.Engine) {

	route := router.Group("/")
	route.GET("_health", func(c *gin.Context) { c.JSON(200, gin.H{"msg": "hello"}) })

	files := route.Group("/file")
	{
		// 1.check file uploaded status
		// 2.check and upload chunk file directly
		files.POST("/upload", UploadFile)

		// 1.merge file with filename, related chunkNum
		// 2.check total chunks match
		// 3.purge tmp files
		files.POST("/merge", MergeFile)
	}
}
