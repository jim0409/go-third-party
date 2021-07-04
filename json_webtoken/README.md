# 簡介
Authentication(認證方式): 允許你的應用程式去了解，這個登入的使用者是否為合法授權。


# quick-test
1. go build

2. login with /signin
   > curl -XPOST "http://localhost:8000/signin" -d '{"username":"user1","password":"password1"}' -v

3. check validation with /welcome
   > curl -XGET "http://localhost:8000/welcome" -v --cookie "token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InVzZXIxIiwiZXhwIjoxNTcwMTIzNDMwfQ.XovVAYPnzlyKHZBgCQiA0tLHGweGk-zrafHkoy1ERa0"

4. refresh token with /refresh
   > curl -XGET "http://localhost:8000/refresh" -v --cookie "token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InVzZXIxIiwiZXhwIjoxNTcwMTIzNDMwfQ.XovVAYPnzlyKHZBgCQiA0tLHGweGk-zrafHkoy1ERa0"

# refer:
- https://www.sohamkamani.com/blog/golang/2019-01-01-jwt-authentication/