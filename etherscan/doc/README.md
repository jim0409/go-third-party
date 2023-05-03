### intro

區塊鏈 API 查詢系統文件檔

### 功能流程
1. 註冊使用者錢包
<!-- 
- CRUD: 支持註冊/讀取/更新/刪除 監聽的錢包 __ simple CRUD to msyql
-->

2. 基於限流速桶, 定期撈取使用者的錢包交易紀錄
<!-- 
- LeakBucket ___  "go.uber.org/ratelimit"
- 根據使用者錢包派送撈取交易訂單事件
- 根據上一次撈取的訂單號完成時間
- 錯誤處理
-->

3. 比對差異的交易紀錄, 做後續訪問請求
<!-- 
- 枚舉可能存在的差異?
- 訪問失敗處置
- 資料修改
-->


### 主要管理物件 & 管理物件介面
- AccountManager
    - AccountFolder
- RequestPoolManager
    - Worker
- ExchangeCenter
    - ExchangeClient






