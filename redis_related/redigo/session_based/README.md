# 使用redis儲存使用者的登入session
當一個使用者登入某一個網頁時，server會儲存該使用者的session並且返回給前端，並且該session可以提供給該使用者下次登入時做參照用。

# quick start
1. get session token (依照使用者產生對應的token)
```
curl -XPOST http://localhost:1234/signin -d '{"username":"user2","password":"password2"}'
```
這個endpt會將使用者的訊息進行驗證，並且針對合法的登入使用者進行session_token的儲存
> Cache.Do("SETEX", sessionToken, "120", creds.Username): 將sessionToken設定為key，creds.Username設定為value

2. use token to check welcome endpoint (使用token去確認對應的使用者資訊)
```
curl -v http://localhost:1234/welcome --cookie "session_token=03ccd630-7588-48cf-8333-094eacd41200"
```
> Cache.Do("GET", sessionToken): 這個endpt會將使用者先前驗證的訊息作儲存，並將session對應的username打印出來

3. refresh token (更新session_token的有效期限)
```
curl -v http://localhost:1234/refresh --cookie "session_token=df1ec034-10db-4d1e-b21c-5cc82d575a7e"
```
> Cache.Do("SETEX", newSessionToken, "120", fmt.Sprintf("%s", response)): 這個endpt會將對應的session_token的到期時間重置為120s

# 定義兩個endpt
- /signin: 使用者可以透過這個endpt傳送登入資訊(名稱/密碼)
- /welcome: 使用者可以透過這個endpt看到自身的登入資訊

# refer:
- https://www.sohamkamani.com/blog/2018/03/25/golang-session-authentication/#overview
- https://mycodesmells.com/post/using-redis-as-sessions-store-in-go