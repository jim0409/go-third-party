package main

import "github.com/gin-gonic/gin"

func OkResp(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(code, map[string]interface{}{
		"meta": map[string]interface{}{
			"code":    code,
			"status":  "success",
			"message": msg,
		},
		"data": data,
	})
}

func ErResp(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(code, map[string]interface{}{
		"meta": map[string]interface{}{
			"code":    code,
			"status":  "failed",
			"message": msg,
		},
		"data": data,
	})
	c.Abort()
}
