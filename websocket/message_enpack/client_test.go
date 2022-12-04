package main

import (
	"log"
	"net/http"
	"testing"

	"github.com/gorilla/websocket"
)

func TestMockSer(t *testing.T) {

	http.Handle("/", http.HandlerFunc(wsHandler))
	http.ListenAndServe(":3000", nil)
	select {}
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
	// time.Sleep(time.Second)
}

func readMsg(c *websocket.Conn) {
	for {
		_, msg, err := c.ReadMessage()
		if err != nil && err != websocket.ErrCloseSent {
			log.Println("read:", err)
			return
		}
		// log.Printf("receive: %s\n", msg)
		c.WriteMessage(websocket.TextMessage, msg)
	}
}
