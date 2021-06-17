# 引言
數據上，假使 msg 格式像這樣
```json
{
    "type": "UPDATE",
    "database": "blog",
    "table": "blog",
    "data": [
        {
            "blogId": "100001",
            "title": "title1",
            "content": "this is a blog",
            "uid": "1000012",
            "state": "1"
        }
    ]
},
{
    "type": "UPDATE",
    "database": "blog",
    "table": "comment",
    "data": [
        {
            "commentId": "100002",
            "title": "title2",
            "comment": "this is a blog",
            "cuid": "1000012",
            "state": "1"
        }
    ]
}
```


## json 轉化為 map
```log
map[data:[map[title:title content:this is a blog uid:1000012 state:1 blogId:100001]] type:UPDATE database:blog table:blog]
```
問題
- 通過key 獲取數據，可能出現不存在的key，為了嚴謹，需要檢查key 是否存在；
- 相對於結構體的方式，map數據提取不便且不能利用IDE 補全檢查，key 容易寫錯；


## json 轉化為 struct
問題
- 部分JSON結構`data`，JSON成功解析前無法提前知道
- 轉化的結構體成員必須是可導出的，所以與成員變量名都是大寫，而與JSON的映射通過`json:"tagName"`的tagName完成
```go
type Event struct {
	Type     string              `json:"type"`
	Database string              `json:"database"`
	Table    string              `json:"table"`
	Data     []map[string]string `json:"data"`
}
```
> 即便如此 `Data` 還是 map[string]string{}


## map 轉化為 struct
GO 沒有內建 map 轉化 strcut 的功能。如果要實現，需要依賴於 GO 的反射機制
> 目前有支持的三方套件 [mapstructure](https://github.com/mitchellh/mapstructure)

- 先定義 map 轉化成 struct 結構
```go
type Blog struct {
	BlogId  string `mapstructure:"blogId"`
	Title   string `mapstructrue:"title"`
	Content string `mapstructure:"content"`
	Uid     string `mapstructure:"uid"`
	State   string `mapstructure:"state"`
}

type Comment struct {
	CommentId  string `mapstructure:"commentId"`
	Title   string `mapstructrue:"title"`
	Comment string `mapstructure:"comment"`
	CUid     string `mapstructure:"cuid"`
	State   string `mapstructure:"state"`
}
```

- 實作代碼
```go
e := Event{}
if err := json.Unmarshal(msg, &e); err != nil {
	panic(err)
}

// 透過 Type 來解析 Blog or Comment
switch e.Table {
case "blog":
    var blogs []Blog
    if err := mapstructure.Decode(e.Data, &blogs); err != nil {
        panic(err)
    }
    fmt.Println(blogs)

case "comment":
    var comment []Comment
    if err := mapstructure.Decode(e.Data, &comment); err != nil {
        panic(err)
    }
    fmt.Println(comment)
}
```


## 弱類型解析
前面用 Decode/ Unmarshal 解析出來的 Uid 及 State 應為 int 而非 string
```go
type Blog struct {
	BlogId  string `mapstructure:"blogId"`
	Title   string `mapstructrue:"title"`
	Content string `mapstructure:"content"`
	Uid     int    `mapstructure:"uid"`
	State   int    `mapstructure:"state"`
}
```
解決方法
1. 使用時進行轉化，比如類型為 int 數據，使用strconv.Atoi 轉化 .. 轉化過程太冗繁 .. 而且有可能出錯
2. 使用 mapstructure 提供的類型 map 轉化 struct 功能 .. 透過弱型別轉化

```go
var blogs []Blog
if err := mapstructure.WeekDecode(e.Data, &blogs);err != nil {
    panic(err)
}

fmt.Println(blogs)
```






# refer:
- https://zhuanlan.zhihu.com/p/66926495
- https://segmentfault.com/a/1190000023442894
- https://pkg.go.dev/github.com/mitchellh/mapstructure?utm_source=godoc
- https://github.com/mitchellh/mapstructure
- https://darjun.github.io/2020/07/29/godailylib/mapstructure/
