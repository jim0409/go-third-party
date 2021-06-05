package main

import (
	"fmt"

	"gopkg.in/mcuadros/go-syslog.v2"
)

func main() {
	channel := make(syslog.LogPartsChannel)
	handler := syslog.NewChannelHandler(channel)

	server := syslog.NewServer()
	server.SetFormat(syslog.RFC3164)
	server.SetHandler(handler)
	server.ListenUDP("0.0.0.0:1515")
	server.Boot()

	go handleChannel(channel)

	server.Wait()
}

func handleChannel(c syslog.LogPartsChannel) {
	for j := range c {
		fmt.Println(j)
		fmt.Println()
		fmt.Println()
		for i, v := range j {
			fmt.Println(i, v)
			fmt.Println()
		}
		// value := j["content"].(map[string]interface{})
		// value := j["content"].(string)
		// strArray := strings.Fields(value)

		// for i, v := range strArray {
		// 	fmt.Println(i, v)
		// }
	}
}
