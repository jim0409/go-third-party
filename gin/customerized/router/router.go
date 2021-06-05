package router

import (
	"github.com/gin-gonic/gin"
	"go-third-party/gin/customerized/controllers"
)

func ApiRouter(r *gin.Engine) {
	authrized := r.Group("/")

	// 用於檢查權限以及是否對於請求的Body做必要的預處理
	// authrized.Use(AuthCheck(), AddLog(), BodyPreProcessing())
	authrized.Use(AuthCheck(), AddLog(), ValidateSignature())
	// authrized.Use(AuthCheck(), ValidateSignature())

	r1 := authrized.Group("/v1")
	{
		r1.GET("g1", controllers.Collector)
		r1.POST("g1", controllers.Collector)
	}
}
