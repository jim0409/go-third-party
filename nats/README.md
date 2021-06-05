# intro
研究nats & nats-streaming (使用wrk來確認效能)
- https://toyo0103.github.io/2019/08/19/nats_streaming/

### NATS
##### quick start
> docker run --name nats-main -p 4222:4222 -p 6222:6222 -p 8222:8222 -d nats

##### client (pub & sub) test
1. go run pub_json.go
2. go run sub_json.go

##### installation[binary]
1. download nats-server
- https://github.com/nats-io/nats-server/releases 


### NATS-STREAMING
##### quick start
> docker-compose up -d
(備註:cluster是吃cluster$.config的設定檔，prometheus也是)

##### client (stan-pub & stan-sub) test
1. 運行stan-sub訂閱(訂閱4222)
> go run stan-sub.go -s 127.0.0.1:4222,127.0.0.1:4221,127.0.0.1:4223 -id test -c test -durable durable_name test 

(option)
> go run stan-sub.go -usr jim -pwd password -s 127.0.0.1:4222,127.0.0.1:4221,127.0.0.1:4223 -id test -c test -durable durable_name test


2. 運行stan-pug推送(發送4221)
> go run stan-pub.go -s 127.0.0.1:4221,127.0.0.1:4222,127.0.0.1:4223 -c "test" "test" "123"

(option)
> go run stan-pub.go -usr jim -pwd password -s 127.0.0.1:4221,127.0.0.1:4222,127.0.0.1:4223 -c "test" "test" "123"


##### installation[binary]
1. download nats-streaming-server
- https://github.com/nats-io/nats-streaming-server/releases


# refer:
- https://github.com/nats-io/nats.go
- https://github.com/nats-io/go-nats-examples
- https://blog.csdn.net/villare/article/details/81029958

# natsproxy
- https://nats.io/blog/natsproxy_project/

# nats(docker & related doc)
- https://hub.docker.com/_/nats

# how to config nats-streaming-cluster
- https://carlosbecker.com/posts/nats-streaming-server-cluster
