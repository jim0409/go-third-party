# 透過自定義包的結構來做TCP Server_Client的連結
### 定義共同的協議包，會放入Data
```go
type Package struct {
	Version        [2]byte // 协议版本，暂定V1
	Length         int16   // 数据部分长度
	Timestamp      int64   // 时间戳
	HostnameLength int16   // 主机名长度
	Hostname       []byte  // 主机名
	TagLength      int16   // 标签长度
	Tag            []byte  // 标签
	Msg            []byte  // 日志数据
}
```

### 定義打包方法
```go
func (p *Package) Pack(writer io.Writer) error {
	var err error
	err = binary.Write(writer, binary.BigEndian, &p.Version)
	err = binary.Write(writer, binary.BigEndian, &p.Length)
	err = binary.Write(writer, binary.BigEndian, &p.Timestamp)
	err = binary.Write(writer, binary.BigEndian, &p.HostnameLength)
	err = binary.Write(writer, binary.BigEndian, &p.Hostname)
	err = binary.Write(writer, binary.BigEndian, &p.TagLength)
	err = binary.Write(writer, binary.BigEndian, &p.Tag)
	err = binary.Write(writer, binary.BigEndian, &p.Msg)
	return err
}
```

### 定義解包方法
```go
func (p *Package) Unpack(reader io.Reader) error {
	var err error
	err = binary.Read(reader, binary.BigEndian, &p.Version)
	err = binary.Read(reader, binary.BigEndian, &p.Length)
	err = binary.Read(reader, binary.BigEndian, &p.Timestamp)
	err = binary.Read(reader, binary.BigEndian, &p.HostnameLength)
	p.Hostname = make([]byte, p.HostnameLength)
	err = binary.Read(reader, binary.BigEndian, &p.Hostname)
	err = binary.Read(reader, binary.BigEndian, &p.TagLength)
	p.Tag = make([]byte, p.TagLength)
	err = binary.Read(reader, binary.BigEndian, &p.Tag)
	p.Msg = make([]byte, p.Length-8-2-p.HostnameLength-2-p.TagLength)
	err = binary.Read(reader, binary.BigEndian, &p.Msg)
	return err
}
```

### 定義黏包方法
黏包需要正確的分割自節流中的數據，常見的方法有以下三種
1. 定長分隔(每個數據包最大為該長度)，缺點是數據不足時會浪費傳輸資源
2. 特定字符分隔(如\r\n)，缺點是如果鄭文中有\r\n等字元就會產生問題
3. 在數據包中添加長度自段(如下方例子)
golang提供了buffio.Scanner來解決黏包問題
```go
scanner := bufio.NewScanner(reader) // reader为实现了io.Reader接口的对象，如net.Conn
scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if !atEOF && data[0] == 'V' { // 由于我们定义的数据包头最开始为两个字节的版本号，所以只有以V开头的数据包才处理
		if len(data) > 4 { // 如果收到的数据>4个字节(2字节版本号+2字节数据包长度)
			length := int16(0)
			binary.Read(bytes.NewReader(data[2:4]), binary.BigEndian, &length) // 读取数据包第3-4字节(int16)=>数据部分长度
			if int(length)+4 <= len(data) { // 如果读取到的数据正文长度+2字节版本号+2字节数据长度不超过读到的数据(实际上就是成功完整的解析出了一个包)
				return int(length) + 4, data[:int(length)+4], nil
			}
		}
	}
	return
})
// 打印接收到的数据包
for scanner.Scan() {
	scannedPack := new(Package)
	scannedPack.Unpack(bytes.NewReader(scanner.Bytes()))
	log.Println(scannedPack)
}
```

# refer:
- https://www.ddhigh.com/2018/03/02/golang-tcp-stick-package.html
- https://www.jishuwen.com/d/24nd