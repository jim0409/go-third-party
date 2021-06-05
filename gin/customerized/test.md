1. 針對語意`參數非必填時、填空欄位須參加簽名` ... 
2. 創建帳號的時候，username跟password怎麼給？(看request只有帶參數Account跟Agent)
3. 是否可以給每一個api的對應curl的指令
以/CreateAccount這隻API來看，目前理解是

string key = "Account=YIHAO&Agent=65&key=ea8ab1992149229498d15258ccbe793a"
> echo $key |md5 (dce2ae94b9984f05b85cb4bb1ef72fea)

string sign = Md5(key) <--- 我求出的值是 dce2ae94b9984f05b85cb4bb1ef72fea

curl -XPOST https://api.playgas.tech/api/V2/Game/CreateAccount -H "Content-Type: application/json" -d '{"Account": "YIHAO", "Agent": 65, "Sign": sign}'


# TestApi
1. 夾帶 user: jim
2. 請求需要符合Json格式
3. 請求格式需要符合
```json
{
    "Param":{ ... },
    "Sign": "string",
}
```


```sh
curl --request POST \
  --url http://127.0.0.1:9090/v1/g1 \
  --header 'content-type: application/json' \
  --header 'user: jim' \
  --data '{
	"Param":{
		"url":"http://jimqaweb.mlytics.ai/cache.txt"
	},
	"Sign":"good"
}'
```