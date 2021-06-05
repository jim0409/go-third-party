package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GroupGroup struct {
	groups []*gin.RouterGroup
}

func NewGroupGroup(groups []*gin.RouterGroup) GroupGroup {
	return GroupGroup{
		groups,
	}
}

func (g *GroupGroup) handle(method string, path string, handler gin.HandlerFunc) {
	for _, group := range g.groups {
		group.Handle(method, path, handler)
	}
}

func main() {
	router := gin.Default()
	v1 := router.Group("/v1")
	v2 := router.Group("/v2")

	g := NewGroupGroup([]*gin.RouterGroup{v1, v2})

	g.handle(http.MethodGet, "hello", sayHello)
	g.handle(http.MethodPost, "goodbye", sayGoodbye)
	router.Run()
}

func sayHello(c *gin.Context) {
	c.JSON(200, "hello")
	return
}

func sayGoodbye(c *gin.Context) {
	c.JSON(200, "good bye")
	return
}
