# 使用google Remote Process Control

# 安裝protoc步驟
1. 下載protoc檔案
   1. https://developers.google.com/protocol-buffers/docs/downloads

2. 下載golang的gRPC套件
   1. https://github.com/golang/protobuf
      1. go install github.com/golang/protobuf/protoc-gen-go

# 快速安裝protoc-gen-go
> sh install_grpc.sh

# refer
https://developers.google.com/protocol-buffers/docs/gotutorial
https://ithelp.ithome.com.tw/articles/10207405?sc=hot

# 整理筆記
https://github.com/jim0409/LinuxIssue/blob/master/OS%E7%92%B0%E5%A2%83%E7%9B%B8%E9%97%9C/%E6%9C%89%E9%97%9CgRPC%E5%85%A9%E4%B8%89%E4%BA%8B.md

# 快速筆記使用gRPC參考
https://github.com/jim0409/grpc-go/tree/master/examples

#### some issue happend while you comiple code with mac osx
please refer:
- https://github.com/grpc/grpc-go/issues/2181
```
solution: cd $GOPATH/go/src/golang.org/x/sys/unix; git checkout 1c9583448a9c3aa0f9a6a5241bf73c0bd8aafded
```
