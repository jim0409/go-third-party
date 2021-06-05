# rfc3164_format.md
### rfc3164.go
定義對應的formater資料結構以及要能夠被回傳的方法
1. RFC3164 struct
2. GetParser(line []byte) LogParser
3. GetSplitFunc() bufio.SplitFunc


### 
1. 
```go
type RFC3164 struct{}
```

2. 回傳一個adapte的資料結構
```go
func (f *RFC3164) GetParser(line []byte) LogParser {
	return &parserWrapper{rfc3164.NewParser(line)}
}
```

3. 
```go
func (f *RFC3164) GetSplitFunc() bufio.SplitFunc {
	return nil
}
```