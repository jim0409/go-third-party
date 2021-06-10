package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"go-third-party/websocket/advacne/ws"
)

var (
	hub = ws.NewHub(nil) //新建一个用户
)

func main() {

	go hub.Run() //开始获取用户中传送的数据

	http.HandleFunc("/ws/conn", func(res http.ResponseWriter, r *http.Request) {
		ws.ServeWs(hub, res, r)
	})
	go beatInterval(3)

	anotherBeatInterval(3)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Panic(err)
	}
}

func beatInterval(t int) {
	for {
		time.Sleep(time.Duration(t) * time.Second)
		hub.Broadcast <- []byte(string("this is heart beat message"))
	}
}

func anotherBeatInterval(t int) {
	fn := func(message []byte, hub *ws.Hub) error {
		// log.Println("message:", string(message))
		hub.Broadcast <- []byte(fmt.Sprintf("this is return message %v", string(message)))
		return nil
	}
	hub.OnMessage = fn
}
