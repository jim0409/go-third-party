package main

import (
	"io"
	"log"
	"net/http"
	"time"

	"go-third-party/websocket/gorilla/slot/ws"
)

var (
	hub = ws.NewHub(nil) //新建一个用户
)

func main() {

	go hub.Run() //开始获取用户中传送的数据

	http.HandleFunc("/ws/conn", func(res http.ResponseWriter, r *http.Request) {
		ws.ServeWs(hub, res, r)
	})

	// a corn beat for ws client
	go beatInterval(30)

	// non-block beat interval
	hub.OnMessage = callbackMsg

	// refer to ./networkServer/http-server/simpleHttpServer/main.go
	http.HandleFunc("/health", func(res http.ResponseWriter, r *http.Request) {
		res.Header().Set("Content-Type", "text/html")
		_, err := io.WriteString(res, "ok")
		res.WriteHeader(http.StatusOK)

		if err != nil {
			log.Fatal(err)
		}
	})

	// since http use the default ListenAndServer, hence add ws.HanderFunc inside
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

func callbackMsg(msg []byte, hub *ws.Hub) error {
	hub.Broadcast <- spin()
	return nil
}

func spin() []byte {
	return []byte("1,2,3;4,5,6;7,8,9;10,11,12;13,14,15")
}
