# 介紹
建立大檔案傳輸的後端 API

# 功能實現
1. 分片上傳
2. 秒傳，斷點續傳
3. 確認合併狀態
4. 合併分片
5. (TODO)提供上傳後下載網址 url
 

# 測試腳本
- upload
```shell
curl -F "myFile=@docker-compose.yml" -H "username: jim" \
http://127.0.0.1:8000/file/upload?filename=docker-compose.yml&md5value=9176b139835b4888ef37776bfdeefab6&chunkorder=1&totalchunks=1
```

- merge
```shell
curl -H "username: jim" \
http://127.0.0.1:8000/file/merge?filename=docker-compose.yml \
-d '{
    "chunk_file_md5": [
        "9176b139835b4888ef37776bfdeefab6"
    ]
}'
``` 

<!-- 
curl -H "username: jim" \
http://127.0.0.1:8000/file/merge?filename=auto.mp4 \
-d '{
    "chunk_file_md5": [
        "eb02a78c7158e3cfeeeb2989c23d0920",
        "f7a9cd4cf188f4737a17fba0b58268ee",
        "0417f368ad3d98f048d609c6b7961bd5",
        "0394186975fbdaadcce19313a3c368dd",
        "6dcf4aea79fb898599ea0b10064654ba",
        "10ddea23cda77b8d1efda93aabc656cd",
        "f51f84bd33a4a8f6c663a6d4d701f248",
        "f10b0690de37e097054ca28e11be4462"
    ]
}'
 -->


# 參照:
### 靈感參照
- https://github.com/threeaccents/mahi
- https://github.com/atulkgupta9/big-file-uploader
- https://github.com/zhuchangwu/large-file-upload
- https://medium.com/akatsuki-taiwan-technology/uploading-large-files-in-golang-without-using-buffer-or-pipe-9b5aafacfc16
- https://freshman.tech/file-upload-golang/
- https://stackoverflow.com/questions/39761910/how-can-you-upload-files-as-a-stream-in-go
- https://prajwol-kc.com.np/2022/03/05/chunk-file-uploading-on-golang/
- https://tw511.com/a/01/32972.html

### 緩存文件
- https://wangbjun.site/2020/coding/golang/file-md5.html
- https://www.kancloud.cn/digest/batu-go/153538
