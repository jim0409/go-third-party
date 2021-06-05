package main

import (
	"log"

	"github.com/gorilla/websocket"
)

func main() {

	c, _, err := websocket.DefaultDialer.Dial("ws://127.0.0.1:3000/ws", nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	err = c.WriteMessage(websocket.TextMessage, []byte("hello message"))
	if err != nil {
		log.Println(err)
		return
	}

	_, msg, err := c.ReadMessage()
	if err != nil {
		log.Println("read:", err)
		return
	}
	log.Printf("receive: %s\n", msg)
}
