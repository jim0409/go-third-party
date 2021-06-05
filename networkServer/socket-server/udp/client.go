package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	res, err := sendUDP("127.0.0.1:8001", "hello")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(res)
	}
}
func sendUDP(addr, msg string) (string, error) {

	conn, _ := net.Dial("udp", addr)

	// send to socket
	_, err := conn.Write([]byte(msg))

	// listen for reply
	bs := make([]byte, 1024)
	conn.SetDeadline(time.Now().Add(3 * time.Second))
	len, err := conn.Read(bs)
	if err != nil {
		return "", err
	} else {
		return string(bs[:len]), err
	}
}
