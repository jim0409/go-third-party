package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/jimweng/networkServer/socket-server/tcp/TCP_customer_protocol/server_client_framework/utils"
)

// Msg 表示擬定要傳送的一個資料結構，該格式擬定如下
type Msg struct {
	Meta    map[string]interface{} `json:"meta"`
	Content interface{}            `json:"content"`
}

func send(conn net.Conn) {
	for i := 0; i < 6; i++ {

		// 做一個模擬判斷，讓不合法的meta data帶入，看是否會如期回傳"Hey! ..."
		var metavalue string
		if i%2 == 0 {
			metavalue = "test"
		} else {
			metavalue = "testNotFound"
		}

		session := GetSession()
		// 定義出client要送給server的資料結構賦值
		message := &Msg{
			Meta: map[string]interface{}{
				"meta": metavalue,
				"ID":   strconv.Itoa(i),
			},
			Content: Msg{
				Meta: map[string]interface{}{
					"author": "JimWeng",
					"age":    13,
					"habbit": "coding",
				},
				Content: session,
			},
		}
		result, _ := json.Marshal(message) // 將資料結構做json化
		conn.Write(utils.Enpack((result))) // 將json格式透過先前定義好的protocol做打包
		//conn.Write([]byte(message))

		time.Sleep(1 * time.Second) // 每次傳送資料錢都做一小段間隔，也可以用來測試heartbeat的超時連線檢測是否生效
	}
	fmt.Println("send over")
	defer conn.Close()
}

// GetSession 每次都會回傳一個擋下的UnixTimestamp用以期待每次送出的session值是迭代的
func GetSession() string {
	gs1 := time.Now().Unix()
	gs2 := strconv.FormatInt(gs1, 10)
	return gs2
}

func main() {
	server := "localhost:1024"
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

	fmt.Println("connect success")

	// 開啟一條goroutine來處理收到的包
	go receive(conn)

	send(conn)

}

// 用來處理收到的包要對應做什麼事
func receive(conn net.Conn) {
	for {
		// deal return session
		tmpBuff := make([]byte, 1024)
		n, err := conn.Read(tmpBuff)
		if err != nil {
			log.Println(err)
			continue
		}
		log.Println("message receive: ", string(tmpBuff[:n]))
	}
}
