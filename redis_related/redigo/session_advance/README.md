# 使用redis當作session紀錄

# 實做流程
0. 使用framework `echo` (可以替換成gin)
1. 定義`Store`作為`Session`儲存用的一個接口
   1. Get("使用者名稱") Session, 錯誤代碼
   2. Set("使用者名稱", Session) 錯誤代碼
2. 設置middleware，讓使用者在通過時可以儲存
3. 使用`curl localhost:5000 --cookie-jar /tmp/cookie-jar --cookie /tmp/cookie-jar`做測試

# 使用Redis來承接Store這個接口


# 考慮

# refer:
- https://mycodesmells.com/post/using-redis-as-sessions-store-in-go