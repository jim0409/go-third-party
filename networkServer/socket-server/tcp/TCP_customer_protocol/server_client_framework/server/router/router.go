package router

import (
	"encoding/json"
	"log"
	"net"
)

type Message struct {
	Meta    map[string]interface{} `json:"meta"`
	Content interface{}            `json:"content"`
}

type Controller interface {
	Excute(Message) []byte
}

type RouterFunc func(Message) bool

type Router struct {
	Routing RouterFunc
	Handler Controller
}

var routers = make([]Router, 0)

func Route(fn RouterFunc, controller Controller) {
	routers = append(routers, Router{fn, controller})
}

func TaskDeliver(postdata []byte, conn net.Conn) {
	var m Message
	err := json.Unmarshal(postdata, &m)
	if err != nil {
		log.Println(err)
		return
	}

	for _, v := range routers {
		do := v.Routing
		act := v.Handler

		if do(m) {
			result := act.Excute(m)
			conn.Write(result)
			return
		}
	}

	conn.Write([]byte("Hey! The meta data is invalid!"))
	return
}
