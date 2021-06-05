package main

import (
	"encoding/json"
	"log"
	"net"
	"time"
)

type Msg struct {
	Content map[string]interface{} `json:"content"`
}

// const UdpAddr = "127.0.01:1515"
const UdpAddr = "127.0.01:514"

// func RFC3164Formatter(p Priority, hostname, tag, content string) string {
// 	timestamp := time.Now().Format(time.Stamp)
// 	msg := fmt.Sprintf("<%d> %s %s %s[%d]: %s",
// 		p, timestamp, hostname, tag, os.Getpid(), content)
// 	return msg
// }

func main() {

	conn, err := net.Dial("udp", UdpAddr)
	if err != nil {
		panic(err)
	}

	for i := 0; i < 6; i++ {
		// var msg
		message := &Msg{
			Content: map[string]interface{}{
				"host": "jim",
			},
		}
		result, _ := json.Marshal(message)

		result = append(result, result...)

		conn.Write(result) // send to socket

		// listen for reply
		bs := make([]byte, 1024)

		conn.SetDeadline(time.Now().Add(3 * time.Second))
		len, err := conn.Read(bs)
		if err != nil {
			panic(err)
		}
		log.Println(string(bs[:len]))
		time.Sleep(time.Second)
	}
}
