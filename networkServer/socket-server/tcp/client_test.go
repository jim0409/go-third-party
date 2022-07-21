package main

import (
	"fmt"
	"net"
	"os"
	"testing"
)

const server = "127.0.0.1:1024"

type Client struct {
	Conn *net.TCPConn
}

func NewTcpClient() *Client {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	return &Client{
		Conn: conn,
	}
}

func (c *Client) Send(bs []byte) {
	c.Conn.Write(bs)
	fmt.Println("send over")
}

func TestSendRawMsg(t *testing.T) {
	client := NewTcpClient()
	client.Send([]byte("hello world!"))

	fmt.Println("connect success")

}
