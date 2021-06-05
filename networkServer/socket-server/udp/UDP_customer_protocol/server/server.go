package main

import (
	"fmt"
	"net"

	"github.com/jimweng/networkServer/socket-server/udp/UDP_customer_protocol/server/router"
	_ "github.com/jimweng/networkServer/socket-server/udp/UDP_customer_protocol/server/router/handler/all"
)

func main() {
	src := "0.0.0.0:8001"
	listener, err := net.ListenPacket("udp", src)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer listener.Close()

	fmt.Printf("UDP server start and listening on %s.\n", src)

	for {
		buf := make([]byte, 1024)
		n, addr, err := listener.ReadFrom(buf)
		if err != nil {
			continue
		}
		go router.TaskDeliver(buf[:n], listener, addr)

	}
}
