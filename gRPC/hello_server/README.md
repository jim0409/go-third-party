# intro

# quick start
1. start grpc Server
> go run server/main.go

2. 使用`grpc client`做測試
> go run client/main.go
```log
➜  hello_server git:(master) ✗ go run client/main.go 
2020/10/09 16:47:18 Greeting: Hello world
```

3. 使用`grpcurl`做為`client`來測試`grpcserver`(沒有使用tls的話，要帶`-plaintext`)
> grpcurl -d '{"name":"jim"}' -plaintext -import-path ./helloworld -proto helloworld.proto localhost:50051 helloworld.Greeter/SayHello
```log
➜  hello_server git:(master) ✗ grpcurl -d '{"name":"jim"}' -plaintext -import-path ./helloworld -proto helloworld.proto localhost:50051 helloworld.Greeter/SayHello

{
  "message": "Hello jim"
}
```

# refer:
- https://blog.yowko.com/grpcurl/
# 使用gRPC創建一個hello server
- 創建指令: `protoc --go_out=plugins=grpc:. *.proto`