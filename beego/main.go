package main

import (
	"github.com/astaxie/beego"
)

type DemoController struct {
	beego.Controller
}

func (c *DemoController) HelloWorld() {
	c.Ctx.WriteString("hello, world")
}

func main() {
	/*
		beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
			AllowAllOrigins: true,
			// AllowOrigins:     []string{"http://192.168.43.52"},
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type", "X-TOKEN"},
			ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
			AllowCredentials: true,
		}))
	*/
	beego.Router("/hello", &DemoController{}, "*:HelloWorld")
	beego.Run(":8000")
}
