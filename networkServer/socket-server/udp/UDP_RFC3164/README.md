# intro

透過開源專案，學習如何利用golang內建的`net.ListenUDP()`來實作UDP協議RF3164

# package 解說
```go
syslog - server.go
       - handler.go

format - format.go
       - rfc3164.go

internal/syslogparser - syslogparser.go
                      - rfc364
                        - rfc3164.go
```

# refer:
- https://github.com/mcuadros/go-syslog
