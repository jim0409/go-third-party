package router

import (
	"encoding/json"
	"log"
	"net"
)

type Msg struct {
	Meta    map[string]interface{} `json:"meta"`
	Content interface{}            `json:"content"`
}

type Controller interface {
	Excute(Msg) []byte
}

type RouterFunc func(Msg) bool

type Router struct {
	Handler RouterFunc
	Action  Controller
}

var routers = make([]Router, 0)

func Route(fn RouterFunc, controller Controller) {
	routers = append(routers, Router{fn, controller})
}

func TaskDeliver(postdata []byte, conn net.Conn) {
	var entermsg Msg
	err := json.Unmarshal(postdata, &entermsg)
	if err != nil {
		log.Println(err)
		return
	}

	for _, v := range routers {
		pred := v.Handler
		act := v.Action

		if pred(entermsg) {
			result := act.Excute(entermsg)
			conn.Write(result)
			return
		}
	}

	conn.Write([]byte("Hey! The meta data is invalid!"))
	return
}
