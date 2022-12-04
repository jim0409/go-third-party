package main

import (
	"context"
	"fmt"
	"go-third-party/websocket/handshake/message"
	"log"
	"time"

	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
)

type EvtType int

func serialize(msg proto.Message) ([]byte, error) {
	data, err := proto.Marshal(msg)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func wrapMessage(path string, msg proto.Message) (*message.Message, error) {
	data, err := serialize(msg)
	if err != nil {
		return nil, err
	}
	return &message.Message{
		Type: message.Notify,
		Data: data,
	}, nil
}

type wsClient struct {
	Addr string
	Conn *websocket.Conn
}

func newWsClient(addr string, ctx context.Context) (*wsClient, error) {
	conn, err := connect(addr, ctx)
	if err != nil {
		return nil, err
	}

	return &wsClient{
		Addr: addr,
		Conn: conn,
	}, nil
}

func connect(addr string, ctx context.Context) (*websocket.Conn, error) {
	c, _, err := websocket.DefaultDialer.DialContext(ctx, addr, nil)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c *wsClient) readpump() {
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

func (c *wsClient) sendmsg() {
	msg := "ping"
	conn := c.Conn
	ticker := time.NewTicker(1 * time.Second)
	for tk := range ticker.C {
		fmt.Println(tk)
		conn.WriteMessage(websocket.TextMessage, []byte(msg))
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	client, err := newWsClient("ws://127.0.0.1:3000", ctx)
	if err != nil {
		panic(err)
	}

	go client.readpump()
	for {
		client.sendmsg()
		time.Sleep(1 * time.Second)
	}

}
