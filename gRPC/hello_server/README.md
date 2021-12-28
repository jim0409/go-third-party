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
- https://grpc.io/docs/languages/go/quickstart/#regenerate-grpc-code

# 使用gRPC創建一個hello server
<!--
- 創建指令: `protoc --go_out=plugins=grpc:. *.proto`

... 2021/12/28 更新 ...

--go_out: protoc-gen-go: plugins are not supported; use 'protoc --go-grpc_out=...' to generate gRPC
See https://grpc.io/docs/languages/go/quickstart/#regenerate-grpc-code for more information.
-->
- 創建 pb.go & _grpc.bp.go 指令:
```
cd ./helloworld
protoc -I . --go-grpc_out=. --go_out=. ./helloworld/*.proto
```

# extend-refer
### 安裝 protoc
1. brew install protoc


### 安裝 protoc-gen-go-grpc
2. go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest


### 設定 GOBIN/PATH:
```
export GOPATH="/Users/jim.weng/go"
export PATH="$PATH:$GOPATH/bin"
```
