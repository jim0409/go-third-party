package controllers

import (
	"github.com/gin-gonic/gin"
	"go-third-party/gin/reverse_proxy_test/service"
)

func Api(c *gin.Context) {
	service.Api1()
	c.JSON(200, "ok")
	return
}
