# 簡介
### 對於實踐基礎的Server_Client Socket連接，需要至少實踐四件事
1. 自定義一個`通訊協議`
2. 通過`心跳`機制維護連接
3. 通過`router-controller`機制解耦服務器
4. 可以配合配置文件，動態調整系統參數

# refer:
- https://www.jianshu.com/p/bb3994fa78dd
- https://github.com/gislu/goSocket


# refer_如何處理close
- https://www.reddit.com/r/golang/comments/5jo972/how_to_handle_a_lot_of_close_wait_tcp_connections/
- https://dotblogs.com.tw/dizzydizzy/2019/02/15/tcpwithgo
- https://github.com/golang/go/issues/10940
- https://thenotexpert.com/golang-tcp-keepalive/