# intro
考慮到叢集redis以及failover，需要使用redis-cluster

# setup test env quickly
> docker run --rm -p 7000:7000 -p 7001:7001 -p 7002:7002 -p 7003:7003 -p 7004:7004 -p 7005:7005 --name redis-cluster-script -e "IP=0.0.0.0" -d grokzen/redis-cluster

# notes:
目前有直接支持redis-cluster的package為go-redis，所以會採用go-redis來實作

# refer:
- https://www.itread01.com/content/1545897373.html

個人實測試
- https://github.com/jim0409/LinuxIssue/tree/master/redis-cluster

網路教學
- https://www.google.com/search?q=docker+run+redis+cluster&oq=docker+run+redis+cluster&aqs=chrome.0.69i59j35i39j69i60.3476j0j7&sourceid=chrome&ie=UTF-8

# redis-lua
- https://www.freecodecamp.org/news/a-quick-guide-to-redis-lua-scripting/
- https://studygolang.com/articles/18031
- https://www.jianshu.com/p/fec18f59ff8f
- https://www.jishuwen.com/d/2ZKf/zh-tw
