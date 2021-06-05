# syslog
### server.go
1. 定義Server資料結構
```go
type Server struct {
    // 定義要實作的listeners
    listeners               []net.Listener
    // 定義要實作的connections，會橋接net.PacketConn
    connections             []net.PacketConn
    // 
    wait                    sync.WaitGroup
    // 監聽伺服器關閉事件
    doneTcp                 chan bool
    // 定義發送數據框的通道
    datagramChannel         chan DatagramMessage
    // *** 定義要針對數據進行格式封裝的方法
    format                  format.Format
    // 
	handler                 Handler
    
    lastError               error
    
    readTimeoutMilliseconds int64
    
    datagramPool            sync.Pool
}
```

2. 宣告初始化伺服器應該有的物件
```go
//NewServer returns a new Server
func NewServer() *Server {
	return &Server{datagramPool: sync.Pool{
		New: func() interface{} {
			return make([]byte, 65536)
		},
	}}
}
```


3. 後置一些對應的物件屬性
- format.Format
- Handler
- readTimeoutMilliseconds
```go
//Sets the syslog format (RFC3164 or RFC5424 or RFC6587)
func (s *Server) SetFormat(f format.Format) {
	s.format = f
}

//Sets the handler, this handler with receive every syslog entry
func (s *Server) SetHandler(handler Handler) {
	s.handler = handler
}

//Sets the connection timeout for TCP connections, in milliseconds
func (s *Server) SetTimeout(millseconds int64) {
	s.readTimeoutMilliseconds = millseconds
}
```


4. 啟用要聽的udp listener
```go
//Configure the server for listen on an UDP addr
func (s *Server) ListenUDP(addr string) error {
	udpAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return err
	}

	connection, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		return err
	}
	connection.SetReadBuffer(datagramReadBufferSize)

	s.connections = append(s.connections, connection)
	return nil
}
```


5. 啟動伺服器
```go
//Starts the server, all the go routines goes to live
func (s *Server) Boot() error {
	if s.format == nil {
		return errors.New("please set a valid format")
	}

	if s.handler == nil {
		return errors.New("please set a valid handler")
	}

	for _, listener := range s.listeners {
		s.goAcceptConnection(listener)
	}

	if len(s.connections) > 0 {
		s.goParseDatagrams()
	}

	for _, connection := range s.connections {
		s.goReceiveDatagrams(connection)
	}

	return nil
}
```

6. 依據上面啟動的三個方法，來實作啟動伺服器後應該做的事
    1. goAcceptConnection(listener) : 透過 listner.Accept()來返回`connection`
    2. goParseDatagrams() : 透過之前的TODO
    3. goReceiveDatagrams(connection)


#### goAcceptConnection(listener net.Listener) { ... }
- 將listener.Accept()啟動後，進行掃描連線`goScanConnection(connection)`
```go
func (s *Server) goAcceptConnection(listener net.Listener) {
	s.wait.Add(1)
	go func(listener net.Listener) {
	loop:
		for {
			select {
			case <-s.doneTcp:
				break loop
			default:
			}
			connection, err := listener.Accept()
			if err != nil {
				continue
			}

			s.goScanConnection(connection)
		}

		s.wait.Done()
	}(listener)
}
```



- 透過listen.Accept()獲得的connection，放進bufio.NewScanner(...)，connection.RemoteAddr()拿到連線的client地址；同時宣告一個新的資料結構`&ScanCloser{scanner, connection}`
```go
type ScanCloser struct {
    *bufio.Scanner
    // note... 直接繼承 buffio.Scanner 的方法 `Split`
    /*
    func (s *Scanner) Split(split SplitFunc) {
        if s.scanCalled {
            panic("Split called after Scan")
        }
        s.split = split
    }
    */
	closer TimeoutCloser  // 實作closer的兩個方法，Close()以及SetReadline(t time.Time) error
}

type TimeoutCloser interface {
	Close() error
	SetReadDeadline(t time.Time) error
}

////////

func (s *Server) goScanConnection(connection net.Conn) {
    scanner := bufio.NewScanner(connection)
    // 將 split function 賦予給scanner
	if sf := s.format.GetSplitFunc(); sf != nil {
		scanner.Split(sf)
	}

    // 拿到 client 的 remote address
	remoteAddr := connection.RemoteAddr()
	var client string
	if remoteAddr != nil {
		client = remoteAddr.String()
	}

    // 備註: 因為這邊不考慮udp server直接暴露的問題，所以把tls相關都移除了
	tlsPeer := ""

	var scanCloser *ScanCloser
	scanCloser = &ScanCloser{scanner, connection}

	s.wait.Add(1)
	go s.scan(scanCloser, client, tlsPeer)
}
```



- 在透過 listener -> connection -> scanner 以後，進行scan，將收到的文本`scanCloser.Text()`做解析
```go
func (s *Server) scan(scanCloser *ScanCloser, client string, tlsPeer string) {
loop:
	for {
		select {
		case <-s.doneTcp:
			break loop
		default:
		}
		if s.readTimeoutMilliseconds > 0 {
			scanCloser.closer.SetReadDeadline(time.Now().Add(time.Duration(s.readTimeoutMilliseconds) * time.Millisecond))
		}
        if scanCloser.Scan() { // 調用scanCloser內buffio.Scanner的方法
            // 因為之前已經將scanClose的split函數賦予了，調用Scan則是執行scan split來操作split
			s.parser([]byte(scanCloser.Text()), client, tlsPeer)
		} else {
			break loop
		}
	}
	scanCloser.closer.Close() // 調用之前定義過的closer的方法

	s.wait.Done()
}
```



- parser ... 在scanCloser的定義下，有一個buffio.Scanner要調用到Scan，而scan出來的文字。透過parser來做進一步拆解
```go
func (s *Server) parser(line []byte, client string, tlsPeer string) {
    // 宣告 parser 是server format.Format下調用的方法GetParser的回傳
	parser := s.format.GetParser(line)
	err := parser.Parse()
	if err != nil {
		s.lastError = err
	}

    // 將 parser 的值 Dump 出來 ... 回傳一個 LogParts 為 map[strint]interface{}
	logParts := parser.Dump()
	logParts["client"] = client
	if logParts["hostname"] == "" && (s.format == RFC3164) {
		if i := strings.Index(client, ":"); i > 1 {
			logParts["hostname"] = client[:i]
		} else {
			logParts["hostname"] = client
		}
	}
	logParts["tls_peer"] = tlsPeer

	s.handler.Handle(logParts, int64(len(line)), err)
}
```



### goParseDatagrams()
宣告一個chan DatagramMessage限定長度datagramChannelBufferSize給Server.datagramChannel

；收集Server datagramChannel 裡面的數據，將數據拉出後透過指定的SplitFunc做parser，

(如果SplitFunc為nil，則使用server預設的parser進行解析)
```go
func (s *Server) goParseDatagrams() {
	s.datagramChannel = make(chan DatagramMessage, datagramChannelBufferSize)

	s.wait.Add(1)
	go func() {
		defer s.wait.Done()
		for {
			select {
			case msg, ok := (<-s.datagramChannel):
				if !ok {
					return
				}
				if sf := s.format.GetSplitFunc(); sf != nil {
					if _, token, err := sf(msg.message, true); err == nil {
						s.parser(token, msg.client, "")
					}
				} else {
					s.parser(msg.message, msg.client, "")
				}
				s.datagramPool.Put(msg.message[:cap(msg.message)])
			}
		}
	}()
}
```



#### goReceiveDatagrams(packetconn net.PacketConn) { ... }
由listener.Acept()回傳的connection，透過connection.ReadFrom(buf)來拿到封包的內容。

(備註: buf可以從s.datagramPool中拿取，循環利用。避免每次都要宣告拿取造成浪費)
```go
func (s *Server) goReceiveDatagrams(packetconn net.PacketConn) {
	s.wait.Add(1)
	go func() {
		defer s.wait.Done()
		for {
			buf := s.datagramPool.Get().([]byte)
			n, addr, err := packetconn.ReadFrom(buf)
			if err == nil {
				// Ignore trailing control characters and NULs
				for ; (n > 0) && (buf[n-1] < 32); n-- {
				}
				if n > 0 {
					var address string
					if addr != nil {
						address = addr.String()
					}
					s.datagramChannel <- DatagramMessage{buf[:n], address}
				}
			} else {
				// there has been an error. Either the server has been killed
				// or may be getting a transitory error due to (e.g.) the
				// interface being shutdown in which case sleep() to avoid busy wait.
                
                // 透過 net.Error來判斷網路封包的錯誤訊息，只要不是 ECONNRESET 或 ECONNABORTED 都視為錯誤並且返回
                opError, ok := err.(*net.OpError)
				if (ok) && !opError.Temporary() && !opError.Timeout() {
					return
				}
				time.Sleep(10 * time.Millisecond)
			}
		}
	}()
}
```

