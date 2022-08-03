package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go-third-party/websocket/protobuf/msg"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
)

var (
	addr     = "ws://127.0.0.1:12100"
	mockAddr = "http://127.0.0.1:8000/regist/514?currency=CNY&sn="
	testSn   = "cp01"
)

func connect(addr string) (*websocket.Conn, error) {
	c, _, err := websocket.DefaultDialer.Dial(addr, nil)
	if err != nil {
		return nil, err
	}

	return c, nil
}

type wsClient struct {
	Addr  string
	Conn  *websocket.Conn
	login bool
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
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
		}
		buf := bytes.NewBuffer(message)
		var header msg.MSG_HEADER
		err = binary.Read(buf, binary.BigEndian, &header)
		if err != nil {
			log.Println("read header:", err)
		}
		buflen := buf.Len()
		var body []byte
		if buflen > 4 {
			body = buf.Next(buflen)
		}
		switch header {
		case msg.MSG_HEADER_S_C_PONG:
			data := &msg.S_C_Pong{}
			if err := proto.Unmarshal(body, data); err != nil {
				log.Println("print:", err)
				return
			}
			log.Printf("receive: %s\n", data)
		default:
		}
	}

}

func (c *wsClient) PingLoop() {
	msgHeader := msg.MSG_HEADER_C_S_PING
	conn := c.Conn
	ticker := time.NewTicker(1 * time.Second)
	for tk := range ticker.C {
		if c.login {
			fmt.Println(tk)
			buf := new(bytes.Buffer)
			binary.Write(buf, binary.BigEndian, &msgHeader)
			pingMsg := &msg.C_S_Ping{
				Timestamp: proto.Uint64(uint64(tk.Unix())),
			}
			data, _ := proto.Marshal(pingMsg)
			buf.Write(data)
			conn.WriteMessage(websocket.TextMessage, buf.Bytes())
		}

	}
}

func newWsClient(addr string) *wsClient {
	return &wsClient{
		Addr: addr,
	}
}

func GenAccount(sn string) (string, error) {
	resp, err := http.Get(fmt.Sprintf("%v%s", mockAddr, sn))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	m := map[string]interface{}{}

	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(buf, &m)
	if err != nil {
		return "", err
	}

	return m["sid"].(string), nil
}

func (c *wsClient) SetLogin() {
	c.login = true
}

func (c *wsClient) DealLogin(sid string) error {
	data := &msg.C_S_TryReg{
		Sid: proto.String(sid),
	}
	sendData, err := proto.Marshal(data)
	if err != nil {
		return err
	}
	buf := &bytes.Buffer{}

	header := msg.MSG_HEADER_C_S_TRY_REG
	binary.Write(buf, binary.BigEndian, &header)
	buf.Write(sendData)

	return c.Conn.WriteMessage(websocket.TextMessage, buf.Bytes())
}

func main() {
	c := newWsClient(addr)
	if err := c.Connect(); err != nil {
		panic(err)
	}
	sid, err := GenAccount(testSn)
	if err != nil {
		panic(err)
	}
	if err := c.DealLogin(sid); err != nil {
		panic(err)
	}
	c.SetLogin()
	go c.PingLoop()
	c.ReadPump()

}
