package router

import (
	"github.com/gin-gonic/gin"
	"go-third-party/gin/reverse_proxy_test/controllers"
)

func ApiRouter(r *gin.Engine) {
	authorized := r.Group("/")
	r1 := authorized.Group("/v1")
	{
		r1.GET("s1", controllers.Api)
	}
}
