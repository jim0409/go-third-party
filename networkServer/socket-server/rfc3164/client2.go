// syslogsrv_client.go
package main

import (
	"bytes"
	"net"
)

// import (
// 	"log"

// 	syslog "github.com/RackSec/srslog"
// )

// func main() {
// 	w, err := syslog.Dial("udp", "127.0.0.1:1515", syslog.LOG_ERR, "http")
// 	if err != nil {
// 		log.Fatal("failed to dial syslog")
// 	}

// 	// userAgent := " Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.116 Safari/537.36"
// 	w.SetHostname("149.56.16.104")
// 	// w.Alert("time_local request_method request_uri status body_bytes_sent http_referer http_x_forwarded_for request_time upstream_response_time " + userAgent)
// 	w.Alert("This is a test context")

// }

const udpAddr = "127.0.0.1:1515"

func main() {
	conn, err := net.Dial("udp", udpAddr)
	if err != nil {
		panic(err)
	}
	sendUDP(conn)
}

func sendUDP(conn net.Conn) error {
	// message := map[string]interface{}{
	// 	"Meta": "test",
	// }
	// result, _ := json.Marshal(message)

	// conn.Write(result)

	result := &bytes.Buffer{}
	result.WriteByte('1')
	result.WriteByte('2')
	result.WriteByte('3')
	result.WriteByte('4')
	result.WriteByte('5')
	// log.Println(result.Bytes())

	conn.Write(result.Bytes())
	// conn.Write([]byte("string"))

	return nil
}

// type Writer struct {
// 	priority int
// 	tag      string
// 	hostname string
// 	network  string
// 	raddr    string
// }
