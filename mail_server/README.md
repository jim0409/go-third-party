## intro

透過 google 的 smtp伺服器 寄送信件

## 相依安裝
1. 下載 go
> https://go.dev/dl/
2. 下載對應的 go 套件
> go mod vendor
3. 複製一份 config.ini 並且依照環境修改參數
> cp config.ini.tmpl config.ini

## 快速啟動
1. 建置
> go build -o app .
2. 運行
> ./app


## 測試
### 必需:
1. authcode: 身份驗證
2. id: 選擇對應的模板
3. mail: 郵件發送地址
### 可選:
1. sub: 信件主旨
 
```bash
curl "http://127.0.0.1:8000/msg/send" -d '{"authcode":"test", "id":1, "sub": "good_subject", "mail": "berserker.01.tw@gmail.com", "data":{"Name":"Jim", "URL":"https://demo.testfire.net"}}'

```

## 如何設定自己的 google smtp 伺服器
- https://www.webdesigntooler.com/google-smtp-send-mail

## 參考:
- https://gist.github.com/jpillora/cb46d183eca0710d909a
- https://medium.com/@dhanushgopinath/sending-html-emails-using-templates-in-golang-9e953ca32f3d
