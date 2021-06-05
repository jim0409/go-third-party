package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/jimweng/networkServer/socket-server/udp/UDP_customer_protocol/utils"
)

const UdpAddr = "127.0.01:8001"

func main() {
	conn, err := net.Dial("udp", UdpAddr)
	// if err := sendUDP(UdpAddr); err != nil {
	if err != nil {
		panic(err)
	}

	go sendUDP(conn)

	go sendAnotherUDP(conn)

	go waitsignal()
	<-done
}

var done = make(chan bool, 1)

func waitsignal() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	sig := <-sigs
	fmt.Println(sig)
	done <- true

}

func sendUDP(conn net.Conn) error {
	// conn, _ := net.Dial("udp", addr)

	for i := 0; i < 6; i++ {
		// var msg
		message := &utils.Msg{
			Meta: map[string]interface{}{
				"meta": "test",
			},
		}
		result, _ := json.Marshal(message)

		conn.Write(result) // send to socket

		// listen for reply
		bs := make([]byte, 1024)
		conn.SetDeadline(time.Now().Add(3 * time.Second))
		len, err := conn.Read(bs)
		if err != nil {
			return err
		}
		log.Println(string(bs[:len]))
		time.Sleep(time.Second)
	}
	return nil
}

func GetSession() string {
	gs1 := time.Now().Unix()
	gs2 := strconv.FormatInt(gs1, 10)
	return gs2
}

func sendAnotherUDP(conn net.Conn) error {
	// conn, _ := net.Dial("udp", addr)

	for i := 0; i < 6; i++ {
		// var msg
		message := &utils.Msg{
			Meta: map[string]interface{}{
				"meta": "test2",
			},
		}
		result, _ := json.Marshal(message)

		conn.Write(result) // send to socket

		// listen for reply
		bs := make([]byte, 1024)
		conn.SetDeadline(time.Now().Add(3 * time.Second))
		len, err := conn.Read(bs)
		if err != nil {
			return err
		}
		log.Println(string(bs[:len]))
		time.Sleep(time.Second)
	}
	return nil
}
