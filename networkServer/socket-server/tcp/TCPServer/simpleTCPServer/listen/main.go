package main

import (
	"log"
	"net"
)

var n = 0

func main() {
	netListen, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		panic(err)
	}
	defer netListen.Close()

	for {
		conn, err := netListen.Accept()
		if err != nil {
			panic(err)
		}
		handleConnection(conn)

		conn.Close()
	}
}

func handleConnection(conn net.Conn) {
	var err error
	bs := make([]byte, 1024)

	for {
		if _, err = conn.Read(bs); err != nil {
			log.Println(err)
			break
		}

		// 對於收到的bs做加工
		bss, n := writeback(bs)

		// 將加工過後的bs送回client端
		if _, err = conn.Write(bss[:n]); err != nil {
			log.Println(err)
			break
		}
	}

}

func writeback(bs []byte) ([]byte, int) {
	prefixStr := "hello~~~, the return is "
	bs = append([]byte(prefixStr), bs...)
	return bs, len(bs)
}
