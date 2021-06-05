# 爬蟲練習
1. 訪問網頁https://golang.google.cn/pkg/
2. 抓取網頁下所有標準庫的名稱/連結(href標籤)/敘述
3. 劃分pkg的層級
4. 開一個mysql的table:`pkg_content`，把資料存進去
5. 使用gRPC創建一個管理pkg的Server，提供使用者輸入pkg回覆對應的pkg連結/文檔內容，以及是否有父級別pkg;有的話就帶上父級別的pkg
6. 使用net/http監聽一個端口，實現對外提供一個http api接口，接口可以接收pkg名稱，返回json格式的pkg信息
7. 接口需要使用jwt確認安全，並且該接口會應用到gRPC client去拿pkg資訊，並返回給呼叫者
8. 使用nginx_pass代理http api接口


# 環境設定
1. 架設mysql與phpadmin: docker-compose up -d
   1. 登入mysql admin: http://localhost:8080, 伺服器:`db`;帳號:`root`;密碼:`example`;資料庫:`(留空)`
   2. 從終端機登入mysql: `docker exec -it mysql mysql -u root -p` ... 密碼 `example`
2. 在`mysql`資料庫下創建一個表`pkg_content`
```sql
create user 'jim' IDENTIFIED by 'password';
create database `pkg_lists`;
grant all privileges on pkg_lists.* to 'jim';
```

# 實作步驟
1. 使用goquery將網頁的資料做解析
2. 使用gorm將解析出來的資料寫入mysql
3. 使用proto產生gPRC伺服器以及用戶所需要的文件，定義出`message`方法拉取資料
4. 


# 參考
- 使用goquery解析html
https://github.com/PuerkitoBio/goquery

- 使用docker-compose快速啟用前置環境
https://docs.docker.com/samples/library/mysql/

- 使用gorm將資料寫入db
https://github.com/jinzhu/gorm

# 其他補充
- https://stackoverflow.com/questions/27933866/use-goquery-to-find-a-class-whose-value-contains-whitespace
- https://studygolang.com/articles/4602

# MySQL check status syntax
> show status like "%connect";
https://stackoverflow.com/questions/7432241/mysql-show-status-active-or-total-connections
- Connections: The number of connection attempts (successful or not) to the MySQL server.
- Threads_connected: The number of currently open connections.
