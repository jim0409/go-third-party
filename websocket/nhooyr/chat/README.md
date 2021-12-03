# intro

透過 nhooyr 的 websocket 做 chatroom，實現廣播封包

# quick start
> go run . localhost:0

# test
> go test

# 前端
1. index.html/ index.js/ index.css 透過 DOM 設計，可以發送、接收使用者消息

# 後端
1. 使用者透過 `/subscribe` 這個端點接收 ws 的訊息
2. 透過 `/publish` 這個端點做 http post 發送訊息


# refer:
- https://github.com/nhooyr/websocket/tree/master/examples/chat
