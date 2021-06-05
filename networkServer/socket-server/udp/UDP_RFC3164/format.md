# format
### format.go
定義要規則化的`接口資料結構`以及`對應的方法`

1. 定義`LogParts`為主要承接的資料結構
```go
type LogParts map[string]interface{}
```

2. 
```go
type LogParser interface {
    Parse() error
    Dump() LogParts // 透過 Dump會回傳 LogParts
    Location(*time.Location)
}
```

3. 
```go
// server.go 內的 Server 結構中的 format.Format
type Format interface {
    GetParser([]byte) LogParser // 透過 GetParser 會回傳 LogParser
    // type SplitFunc func(data []byte, atEOF bool) (advance int, token []byte, err error)
    GetSplitFunc() bufio.SplitFunc // 透過 GetSplitFunc 拿到同 bufio.SplitFunc 的 interface
}
```

4. 
```go
type parserWrapper struct {
    syslogParser.LogParser
}
```

5. 
```go
func (w *parserWrapper) Dump() LogParts {
    return LogParts(w.LogParser.Dump())
}
```