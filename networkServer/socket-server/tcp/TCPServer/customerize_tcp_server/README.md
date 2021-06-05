# 如何調度這個包

```go
package main

import "github.com/firstrow/tcp_server"

func main() {
	server := tcp_server.New("localhost:9999")

	server.OnNewClient(func(c *tcp_server.Client) {
		// new client connected
		// lets send some message
		c.Send("Hello")
	})
	server.OnNewMessage(func(c *tcp_server.Client, message string) {
		// new message received
	})
	server.OnClientConnectionClosed(func(c *tcp_server.Client, err error) {
		// connection with client lost
	})

	server.Listen()
}
```

# 這是一個別人的專案，放入這邊只是方便以後查找，以及練習寫一些單元測試，沒有要向下開發的意思...


# refer:
- https://github.com/firstrow/tcp_server