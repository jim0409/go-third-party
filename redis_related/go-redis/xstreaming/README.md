# intro
streaming 是 redis 在 5.0 後支持的一個 persistent queue 方法

Redis Stream 主要用於消息隊列(MQ, Message Queue)

Redis 本身是有一個 Redis 發佈訂閱 (Pub/Sub) 的機制來實現消息隊列的功能

但他的缺點是消息無法持久化。如果網路、Redis Server 故障，皆有可能導致訊息被丟棄

Streaming 提供了資料持久化 及 分割備份的功能，確保訊息不掉失

- 架構如下
![image](./redis-xstreaming.png)

1. 每個 Stream 都有唯一的名稱，即 Redis 的 Key; 在首次使用 `xadd` 指令創建消息後會自動創建
2. Consumer Group: 消費群，使用`XGROUP CREATE`命令創建，一個消費群可以有多個消費者
3. last_delivered_id: 游標，每個消費群都會有個`last_delivered_id`，對任意一個消費者讀取了消息都會使游標往前移動
4. pending_ids: 消費者(Consumer)的狀態變量，作用是維護消費者的未確認的id;(pending_ids 紀錄當前已經被客戶讀取的消息，但是還沒有 ack (確認)

# 消息隊列相關指令
1. XADD - 添加消息到末尾
2. XTRIM - 對流進行修剪，限制長度
3. XDEL - 刪除消息
4. XLEN - 獲取流包含的元素數量，即消息長度
5. XRANGE - 獲取消息列表，會自動過濾已經刪除的消息
6. XREVRANGE - 反向獲取消息列表，ID從大到小
7. XREAD - 以阻塞或非阻塞的方式獲取消息列表

# 消費者群相關命令:
1. XGROUP CREAE - 創建消費者組
2. XREADGROUP GROUP - 讀取消費者組中的消息
3. XACK - 將消息標記為 "已處理"
4. XGROUP SETID - 為消費者組設置新的最後遞送消息 ID
5. XGROUP DELCONSUMER - 刪除消費者
6. XGROUP DESTORY - 刪除消費者組
7. XPENDING - 顯示待處理消息的相關信息
8. XCLAIM - 轉移消息的歸屬權
9. XINFO - 查看流和消費者組的相關信息
10. XINFO GROUPS - 打印消費者組的信息
11. XINFO STREAM - 打印流信息


### XADD
- 使用`XADD`向隊列添加消息，如果指定的隊列不存在，則創建一個隊列
> XADD key ID field [field value ...]
1. key        : 列隊名稱，如果不存在就創建
2. ID         : 消息 id，使用 * 表示由 redis 生成，可以自定義。但是要自己保證遞增性
3. field value: 紀錄
```
redis> XADD mystream * name Sara surname OConnor
"1601372323627-0"
redis> XADD mystream * field1 value1 field2 value2 field3 value3
"1601372323627-1"
redis> XLEN mystream
(integer) 2
redis> XRANGE mystream - +
1) 1) "1601372323627-0"
   2) 1) "name"
      2) "Sara"
      3) "surname"
      4) "OConnor"
2) 1) "1601372323627-1"
   2) 1) "field1"
      2) "value1"
      3) "field2"
      4) "value2"
      5) "field3"
      6) "value3"
redis>
```

### XTRIM
- 使用 XTRIM 對流進行修剪，限制長度
> XTRIM key MAXLEN [~] count
1. key   : 隊列名稱
2. MAXLEN: 長度
3. count : 數量
```
127.0.0.1:6379> XADD mystream * field1 A field2 B field3 C field4 D
"1601372434568-0"
127.0.0.1:6379> XTRIM mystream MAXLEN 2
(integer) 0
127.0.0.1:6379> XRANGE mystream - +
1) 1) "1601372434568-0"
   2) 1) "field1"
      2) "A"
      3) "field2"
      4) "B"
      5) "field3"
      6) "C"
      7) "field4"
      8) "D"
127.0.0.1:6379>

redis>
```

### XDEL
- 使用 XDEL 刪除消息
> XDEL key ID [ID ...]
1. key : 隊列名稱
2. ID  : 消息 ID
```
> XADD mystream * a 1
1538561698944-0
> XADD mystream * b 2
1538561700640-0
> XADD mystream * c 3
1538561701744-0
> XDEL mystream 1538561700640-0
(integer) 1
127.0.0.1:6379> XRANGE mystream - +
1) 1) 1538561698944-0
   2) 1) "a"
      2) "1"
2) 1) 1538561701744-0
   2) 1) "c"
      2) "3"
```

### XLEN
- 使用 XLEN 獲取流包含的元素數量，即消息長度
> XLEN key
1. key: 隊列名稱
```
redis> XADD mystream * item 1
"1601372563177-0"
redis> XADD mystream * item 2
"1601372563178-0"
redis> XADD mystream * item 3
"1601372563178-1"
redis> XLEN mystream
(integer) 3
redis>
```

### XRANGE
- 使用 XRANGE 獲取消息列表，會自動過濾已經刪除的消息
> XRANGE key start end [COUNT count]
1. key  : 隊列名
2. start: 開始值, `-`表示最小值
3. end  : 結束值, `+`表示最大值
4. count: 數量
```
redis> XADD writers * name Virginia surname Woolf
"1601372577811-0"
redis> XADD writers * name Jane surname Austen
"1601372577811-1"
redis> XADD writers * name Toni surname Morrison
"1601372577811-2"
redis> XADD writers * name Agatha surname Christie
"1601372577812-0"
redis> XADD writers * name Ngozi surname Adichie
"1601372577812-1"
redis> XLEN writers
(integer) 5
redis> XRANGE writers - + COUNT 2
1) 1) "1601372577811-0"
   2) 1) "name"
      2) "Virginia"
      3) "surname"
      4) "Woolf"
2) 1) "1601372577811-1"
   2) 1) "name"
      2) "Jane"
      3) "surname"
      4) "Austen"
redis>
```

### XREVRANGE
- 使用 XREVRANGE 獲取消息列表，會自動過濾已經刪除的消息
> XREVRANGE key end start [COUNT count]
1. key  : 隊列名
2. end  : 結束值, `+`表示最大值
3. start: 開始值, `-`表示最小值
4. count: 數量
```
redis> XADD writers * name Virginia surname Woolf
"1601372731458-0"
redis> XADD writers * name Jane surname Austen
"1601372731459-0"
redis> XADD writers * name Toni surname Morrison
"1601372731459-1"
redis> XADD writers * name Agatha surname Christie
"1601372731459-2"
redis> XADD writers * name Ngozi surname Adichie
"1601372731459-3"
redis> XLEN writers
(integer) 5
redis> XREVRANGE writers + - COUNT 1
1) 1) "1601372731459-3"
   2) 1) "name"
      2) "Ngozi"
      3) "surname"
      4) "Adichie"
redis>
```

### XREAD
- 使用 XREAD 以阻塞或非阻塞方式獲取消息列表
> XREAD [COUNT count] [BLOCK milliseconds] STREAMS key [key ...] id [id ...]
1. count        : 數量
2. milliseconds : 可選, 阻塞毫秒數, 沒有設置就是非阻塞模式
3. key          : 隊列名
4. id           : 消息 ID
```
# 從 Stream 頭部讀取兩條消息
> XREAD COUNT 2 STREAMS mystream writers 0-0 0-0
1) 1) "mystream"
   2) 1) 1) 1526984818136-0
         2) 1) "duration"
            2) "1532"
            3) "event-id"
            4) "5"
            5) "user-id"
            6) "7782813"
      2) 1) 1526999352406-0
         2) 1) "duration"
            2) "812"
            3) "event-id"
            4) "9"
            5) "user-id"
            6) "388234"
2) 1) "writers"
   2) 1) 1) 1526985676425-0
         2) 1) "name"
            2) "Virginia"
            3) "surname"
            4) "Woolf"
      2) 1) 1526985685298-0
         2) 1) "name"
            2) "Jane"
            3) "surname"
            4) "Austen"
```


### XGROUP CREATE
使用 XGROUP CREATE 創建消費者組
> XGROUP [CREATE key groupname id-or-$] [SETID key groupname id-or-$] [DESTORY key groupname] [DELCONSUMER key groupname consumername]
1. key       : 隊列名稱，如果不存在就創建
2. groupname : 組名
3. $         : 表示從尾部開始消費，只接受新消息，當前 Stream 消息會全部忽略

- 從頭部開始消費:
> XGROUP CREATE mystream consumer-group-name 0-0

- 從尾部開始消費:
> XGROUP CREATE mystream consumer-group-name $

### XREADGROUP GROUP
使用 XREADGROUP GROUP 讀取消費組中的消息
> XREADGROUP GROUP group consumer [COUNT count] [BLOCK milliseconds] [NOACK] STREAMS key [key ...] ID [ID ...]
1. group        : 消費組名
2. consumer     : 消費者名
3. count        : 讀取數量
4. milliseconds : 阻塞毫秒數
5. key          : 隊列名
6. ID           : 消息 ID
```
XREADGROUP GROUP consumer-group-name consumer-name COUNT 1 STREAMS mystream >
```


# refer:
- https://redis.io/commands/xread
- https://www.runoob.com/redis/redis-stream.html
- https://github.com/go-redis/redis/issues/863
