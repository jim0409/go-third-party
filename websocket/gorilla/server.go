package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	http.Handle("/ws/conn", http.HandlerFunc(wsHandler))
	http.ListenAndServe(":3000", nil)
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	// websocket的Upgrade提供一個Conn: 及其方法 (c *Conn) ...
	c, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
	if err != nil {
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
	}
	go echo(c)
	go readMsg(c)
}

func echo(c *websocket.Conn) {
	if err := c.WriteJSON("hello world"); err != nil {
		log.Println(err)
	}
	time.Sleep(time.Second)
}

func readMsg(c *websocket.Conn) {
	_, msg, err := c.ReadMessage()
	if err != nil && err != websocket.ErrCloseSent {
		// log.Println("read:", err)
		return
	}
	log.Printf("receive: %s\n", msg)
}
