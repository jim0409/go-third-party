package main

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

var addr = "ws://127.0.0.1:3000/ws/conn"

func connect(addr string) (*websocket.Conn, error) {
	c, _, err := websocket.DefaultDialer.Dial(addr, nil)
	if err != nil {
		return nil, err
	}

	return c, nil
}

type Client interface {
	Connect() error
	ReadPump()
	PingLoop()
}

type wsClient struct {
	Addr string
	Conn *websocket.Conn
}

func (c *wsClient) Connect() error {
	wsc, err := connect(c.Addr)
	if err != nil {
		return err
	}
	c.Conn = wsc
	return nil
}

func (c *wsClient) ReadPump() {
	conn := c.Conn
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
		}
		log.Printf("receive: %s\n", msg)
	}

}

func (c *wsClient) PingLoop() {
	msg := "ping"
	conn := c.Conn
	ticker := time.NewTicker(1 * time.Second)
	for tk := range ticker.C {
		fmt.Println(tk)
		conn.WriteMessage(websocket.TextMessage, []byte(msg))
	}
}

func newWsClient(addr string) Client {
	return &wsClient{
		Addr: addr,
	}
}

func TestConnect(t *testing.T) {
	var wsClients []Client

	for i := 0; i < 1; i++ {
		c := newWsClient(addr)
		if err := c.Connect(); err != nil {
			log.Fatal(err)
		}
		wsClients = append(wsClients, c)
	}

	for _, j := range wsClients {
		go j.PingLoop()
	}

	for _, j := range wsClients {
		j.ReadPump()
	}
}
