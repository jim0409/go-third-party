# quick start
1. docker-compose up -d :快速帶起一個反向代理的服務系統
2. 嘗試用curl指令去做直接請求
> curl http://127.0.0.1:10/whoyare
3. 嘗試用curl指令對privoxy做請求
> http_proxy=http://127.0.0.1:8118 curl http://172.25.0.2:10/whoyare

# 備註:
1. 172.25.0.2 應該換成 whoyare的container ip


# refer:
如何使用golang的觀念
- https://gianarb.it/blog/golang-forwarding-proxy
