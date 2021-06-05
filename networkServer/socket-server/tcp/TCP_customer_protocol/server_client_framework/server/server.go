package main

import (
	"log"
	"net"
	"time"

	"github.com/jimweng/networkServer/socket-server/tcp/TCP_customer_protocol/server_client_framework/server/router"
	"github.com/jimweng/networkServer/socket-server/tcp/TCP_customer_protocol/server_client_framework/utils"
)

var (
	host              = "127.0.0.1:1024"
	heartBeatInterval = 5
)

func main() {
	startServer()
}

func startServer() {
	// 透過net包，帶起一個tcp伺服器
	netListen, err := net.Listen("tcp", host)
	if err != nil {
		panic(err)
	}
	defer netListen.Close()

	for {
		/*
			Accept()方法調用該伺服器的接收方法，回傳連結(conn)以及錯誤(err)
			透過conn來`讀取`或是`寫入`資料
		*/
		conn, err := netListen.Accept()
		if err != nil {
			continue
		}
		/*
			透過handleConnection來處理連線
		*/
		go handleConnection(conn, heartBeatInterval)

	}
}

func handleConnection(conn net.Conn, timeout int) {
	// tmpBuffer 會用來做解析包，以及確認conn健康用
	tmpBuffer := make([]byte, 0)

	// buffer用來承接`連結(conn)`讀取的資料
	buffer := make([]byte, 1024)
	messnager := make(chan byte)
	for {
		n, err := conn.Read(buffer) // 使用buffer來承接`連結(conn)`讀取出來的資料以1024 byte矩陣的方式紀錄，並且回傳buffer的長度(n)
		if err != nil {
			switch err.Error() {
			case "EOF":
				log.Println("End of the connection, disconnect from remote addr: ", conn.RemoteAddr().String())
				return
			default:
				log.Println("connection error:", err)
				return
			}
		}

		// 透過Depack，解析自定義的TCP封包，並且將封包內容傳達到tmpBuffer
		tmpBuffer = utils.Depack(append(tmpBuffer, buffer[:n]...))

		// 將解析好的封包以及conn(連線)傳遞到Task
		router.TaskDeliver(tmpBuffer, conn)

		/*
			至此已經完成基本的TCP連結，可以正常做自定義的tcp連線交流。
			但是仍舊有一個問題，網路封包不一定完全可靠。所以要建立一個健康檢查機制
			- 旨在如果建立起來的tcp連線超過一定時間仍舊沒有收到封包，則自動斷連
		*/

		//start heartbeating : 定時對連線(conn)做檢測，並且設定一定的連線時長限制，超過時間則通知斷連
		/*
			建立一個heartBeating機制，只要收到值，就會打印`get message ....`
			並且會再重新設定一個timer，做防超時檢查

			另一部分，如果連線建立過久，並沒有持續收到封包。就進行斷連
		*/
		go HeartBeating(conn, messnager, timeout)

		//check if get message from client : 定義一個不斷去拿tmpBuffer的人
		/*
			當有資料進來的時候，tmpBuffer會有值，此時會把tmpBuffer的值寫到messager這個chan裡面
			當message chan裡面有值的時候，會觸發HeartBeating，從readerChannel收到來自client的msg
			並且重新設定一個timer，nowTime + heart_beat_timeout
		*/
		go GravelChannel(tmpBuffer, messnager)

	}
	defer conn.Close()

}

//HeartBeating, determine if client send a message within set time by GravelChannel
func HeartBeating(conn net.Conn, readerChannel chan byte, timeout int) {
	select {
	case _ = <-readerChannel:
		log.Println(conn.RemoteAddr().String(), "get message, keeping heartbeating...")
		conn.SetDeadline(time.Now().Add(time.Duration(timeout) * time.Second))
		break
	case <-time.After(time.Second * 5):
		log.Println("It's really weird to get Nothing!!!")
		conn.Close()
	}

}

func GravelChannel(n []byte, mess chan byte) {
	for _, v := range n {
		mess <- v
	}
	close(mess)
}
