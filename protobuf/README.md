# intro
Protobuf 是由 Google 開發的一種可跨平台、跨語言的數據交換格式

是一種將結構化資料`序列化(變成二進制)`的方法。資料要比`json`格式更小更輕便

Protobuf 是 `Protocol Buffers` 的簡寫

- Protocol: 協定、協議
```
你我倆約定、協商好、談妥的東西，或者叫做條款
ex: 你給我1000元新台幣，我去幫你到店裡買一包七星菸抽其中一根給你
```
> 可以是A、B服務之間傳遞的格式、交換、處理的事情

- Buffer: 一塊(特定大小的)空間
- Protocol Buffers: 你我講好的協議所用到的某一空間

# 為什麼要用 protobuf
- 優: 跟 `Json` 相比 `protobuf` 性能更高，更加規範
	1. 編碼速度快，數據體積小
	2. 使用統一的規範，不用擔心大小寫不同導致解析失敗的問題

- 缺: 失去便利性
	1. 改動協議字段，需要重新生成文件
	2. 數據沒有可讀性



# Protobuf 的格式
Protobuf 的格式為 `proto`，目前有`proto2`、`proto3`兩種版本，主要為 proto3 閹割掉了一些較不嚴謹的功能
> 主要用法還是會以 `proto3` 為主


- file: school.proto
```
syntax = "proto3";
option go_package = ".;school";

message Teacher{
  string name = 1;
  int32 age = 2;
}
```

# 數字代表的意義
上方的數字 `name=1`、`age=2`的數字1與2，不是賦值的意思(不是說 名字=1、年齡=2)
而是編號、唯一識別瑪，好讓程式識別這個變數
<!-- 因為時候大家都被壓縮成二進制，認不出誰是誰，有編號要認人比較方便 -->

在同一個`message`裏面識別碼不可重複，但不同`message`之間重複就沒關係了

這個識別碼編號方式也沒有什麼硬性規定，通常會由上往下從1、2、3...開始依序給

值得一提的是，編號`1~15`識別碼的區域會使用`1 byte`來做編碼，位於`16~2047`之間的識別碼區域則會用`2 bytes`來做編碼

所以可以將較常使用到的欄位盡量都放在`1~15`的位置，進而減少資料的傳輸量。


# 安裝 protoc
1. brew install protoc


# 安裝 protoc-gen-go-grpc
2. go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest


# 設定 GOBIN/PATH:
```
export GOPATH="/Users/jim.weng/go"
export PATH="$PATH:$GOPATH/bin"
```

# 快速建置 base on grpc 的 protobuf(更新1.3.2之後的版本編譯方式)
1. .proto 檔內需要添加 option
```
option go_package = "/proto";
```
2. 指令更改為
```
protoc -I . --go-grpc_out=. --go_out=. ./proto/*.proto
```
<!-- 備註: 請在相對路徑執行!! -->


# refer:
- https://yami.io/protobuf/
- https://developers.google.com/protocol-buffers/docs/gotutorial
- https://segmentfault.com/a/1190000009277748

# debug-error:
- https://blog.csdn.net/weixin_43851310/article/details/115431651
- https://stackoverflow.com/questions/60578892/protoc-gen-go-grpc-program-not-found-or-is-not-executable