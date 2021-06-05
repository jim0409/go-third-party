# quick start
1. 啟動伺服器
> go run main.go

2. 發起一般請求確認回應
> curl localhost:8080/;curl localhost:8080/1; curl localhost:8080/2
```response
okayokayokay%
```

3. 使用`/add`這個端點請求，排查api行為是否改變
> curl localhost:8080/add
```response
2020/02/22 19:59:26 server started
2020/02/22 20:00:27 add a handler <--- 可以看到增加一個handler改變
```
4. 再次確認指令
> curl localhost:8080/;curl localhost:8080/1; curl localhost:8080/2
```response
okayNotokayokay%
```
> 可以發現?動態改變了路由的回應了!!

# 原因
http請求進入後，會採用callback的方式去呼叫，所以當呼叫`/add`的時候。會動態去替換掉既有的handler
