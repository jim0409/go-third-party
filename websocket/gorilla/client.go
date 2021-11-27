package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

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
}

type wsClient struct {
	Addr string
	Conn *websocket.Conn
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

func (c *wsClient) Connect() error {
	wsc, err := connect(c.Addr)
	if err != nil {
		return err
	}
	c.Conn = wsc
	return nil
}

func newWsClient(addr string) Client {
	return &wsClient{
		Addr: addr,
	}
}

func main() {
	var wsClients []Client

	for i := 0; i < 10; i++ {
		c := newWsClient(addr)
		if err := c.Connect(); err != nil {
			log.Fatal(err)
		}
		wsClients = append(wsClients, c)
	}

	for _, j := range wsClients {
		go j.ReadPump()
	}

	gracefulShutdown()

}

func gracefulShutdown() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}
