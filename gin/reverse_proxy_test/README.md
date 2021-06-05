# intro
最近在業務上，常常會需要使用到一些gRPC的壓力測試。

特別寫了這個框架來承接gRPC client

# flow
1. 使用 gin server 來做 gRPC 轉發
2. 因為單一個gin server只能乘載最多30000筆轉發，在gin server前面掛載nginx做輪詢以提高後面測試壓力
3. 使用工具 wrk 進行 http 請求，同時保留可以客製化gRPC的空間

# build script
1. build api_server
>  GO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api_server -i main.go

2. build grpc_server
> GO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o grpc_server -i main.go

# run with sacle api=n
> docker-compose up --scale api=3 up -d


# refer:
- https://github.com/gin-gonic/gin