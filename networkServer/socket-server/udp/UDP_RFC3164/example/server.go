package main

import (
	"fmt"

	syslog "github.com/jimweng/networkServer/socket-server/udp/UDP_RFC3164"
)

func main() {
	channel := make(syslog.LogPartsChannel)
	handler := syslog.NewChannelHandler(channel)

	server := syslog.NewServer()
	server.SetFormat(syslog.RFC3164)
	server.SetHandler(handler)
	// 可以加入多個listener
	server.ListenUDP("0.0.0.0:514")
	// server.ListenUDP("0.0.0.0:515")

	server.Boot()

	go func(channel syslog.LogPartsChannel) {
		for logParts := range channel {
			fmt.Println(logParts)
		}
	}(channel)

	server.Wait()
}
