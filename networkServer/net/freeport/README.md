# intro

使用場景，當需要隨機延伸出一個 httpServer時，需要取得當下主機能夠使用的tcp port

# example
調度`GetFreePort()`方法可以拿到一個沒有佔據的port

```go
port, err := GetFreePort()
if err != nil {
    log.Fatal(err)
}

fmt.Println(strconv.Itoa(port))
```


# refer:
- https://github.com/phayes/freeport

# extend-refer:
有關net包裡面的 `ListenTCP`

```go
// ListenTCP acts like Listen for TCP networks.
//
// The network must be a TCP network name; see func Dial for details.
//
// If the IP field of laddr is nil or an unspecified IP address,
// ListenTCP listens on all available unicast and anycast IP addresses
// of the local system.
// If the Port field of laddr is 0, a port number is automatically
// chosen.
```