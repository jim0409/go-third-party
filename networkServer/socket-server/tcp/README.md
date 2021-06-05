# Socket-Client framework
### Server Side
1. netListen, err := net.Listen("tcp","localhost:1024") : 相當於bind port, bind address 以及指定使用的協議(protocol)。
2. conn, err := netListen.Accept() : 對於所有收到的packets進行處理，相當於recv-Q
3. n, err := conn.Read(buffer) : 依照buffer size分片讀取connection中的資料


### Client side
1. tcpAddr, err := net.ResolveTCPAddr("tcp4", server) : 向server端發起請求，並且回傳連線位址
2. conn, err := net.DialTCP("tcp", nil, tcpAddr) : 用拿到的連線地址進行撥打，回傳連線(connection)
3. 定義sender，將words轉成`[]byte`，透過`conn.Write()`將資料送出

# refer:
- https://blog.csdn.net/ahlxt123/article/details/47320161
- https://blog.csdn.net/u010824081/article/details/78025662