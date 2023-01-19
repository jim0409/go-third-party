# intro

練習撰寫 FB 社團相關的 api


# 功能介紹
1. 創建社團
2. 修改社團資訊
3. 加入社團
4. 邀請成員加入社團
5. 將成員踢出社團
6. 社團發文
7. 修改社團發文
8. 分享社團連結


# 資料表定義
<!-- 以 gorm model 代表一張 table -->
1. members
```go
type Member struct {
	ID         int     `gorm:"primaryKey;autoIncrement;"`
	Nickname   string  `gorm:"unique;type:varchar(32);comment:使用者名稱"`
	Identify   string  `gorm:"type:varchar(32);comment:bcrypt對照"`
	GroupInfos []Group `gorm:"many2many:group_members;"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt
}
```

2. groups
```go
type Group struct {
	ID        int      `gorm:"primaryKey;autoIncrement;"`
	Name      string   `gorm:"unique;type:varchar(32);comment:群組名稱"`
	Owner     string   `gorm:"type:varchar(32);comment:群組創建者"`
	Members   []Member `gorm:"many2many:group_members;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
```

3. group_members 
```go
type GroupMember struct {
	GroupID  int    `gorm:"primaryKey"`
	MemberID int    `gorm:"primaryKey"`
	Role     string `gorm:"type:tinyint(8);comment:腳色"`
}
```

4. posts
```go
type Post struct {
	ID        int    `gorm:"primaryKey;autoIncrement;"`
	Owner     string `gorm:"type:varchar(32);comment:發文者"`
	Editor    string `gorm:"type:varchar(32);comment:編輯者"`
	Content   string `gorm:"type:varchar(255);comment:內容"`
	LikeNum   int    `gorm:"type:tinyint(8);comment:喜歡數"`
	HeartNum  int    `gorm:"type:tinyint(8);comment:愛心數"`
	ShareUrl  string `gorm:"type:varchar(255);comment:網址"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
```

# 實現功能

- 會員
```go
type IMember interface {
	NewMember(name string, idfy string) error
	AddMembersToGroup(id int, list []int) error
}
```

- 社團
```go
type IGroup interface {
	NewGroup(owner string, name string) error
}
```

- 貼文
```go
type IPost interface {
	NewPost(int, *Post) error
	UpdatePost(int, *Post) error
}


```
