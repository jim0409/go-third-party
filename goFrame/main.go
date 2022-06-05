package main

import (
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func Hello1(r *ghttp.Request) {
	var User struct {
		LoginTime string `json:"loginTime"`
		Name      string `json:"name"`
	}
	User.LoginTime = time.Now().String()
	User.Name = "jim"
	r.Response.WriteJson(User)
}

func Hello2(r *ghttp.Request) {
	r.Response.Write("localhost: Hello2!")
}

func main() {
	s := g.Server()
	// for multi-domain and multi-ports
	s.Domain("127.0.0.1").BindHandler("/", Hello1)
	s.Domain("localhost").BindHandler("/", Hello2)
	s.SetPort(8080, 8081)
	s.Run()
}
