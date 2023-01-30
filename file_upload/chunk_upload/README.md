# intro

# feature
1. 分片上傳
2. 秒傳，斷點續傳
3. 確認合併狀態
4. 合併分片
5. 提供上傳後下載網址 url
 



# test script
- curl
```shell
curl -F "myFile=@docker-compose.yml" -H "username: jim" \
http://127.0.0.1:8000/file/upload?filename=docker-compose.yml&md5value=9176b139835b4888ef37776bfdeefab6
```

- wrk
> 



# refer:
- https://wangbjun.site/2020/coding/golang/file-md5.html
- https://www.kancloud.cn/digest/batu-go/153538