# intro

練習撰寫 FB 社團相關的 api


# 功能介紹
1. 創建會員 .. <IMember>

1. 創建社團 .. <IGroup>
2. 修改社團資訊 .. <IGroup>
3. 加入社團 .. <IGroup>
4. 等待加入社團人員 .. <IGroup>
5. 邀請成員加入社團 .. <IGroup>
6. 將成員踢出社團 .. <IGroup>
7. 分享社團連結 .. <IPost>

1. 社團發文 .. <IPost>
2. 修改社團發文 .. <IPost>
3. 分享社團發文連結 .. <IPost>


# 實現功能

- 會員
```go
type IMember interface {
	NewMember(name string, idfy string) error
}
```

- 社團
```go
type IGroup interface {
	NewGroup(userId int, name string) error
	UpdateGroup(usrId int, gp *Group) error
	ApplyForGroup(groupId int, userId int) error
	AwaitJoinGroupList(groupId int) ([]int, error)
	AddMembersToGroup(groupId int, memberList []int) error
	KickMemberOutOfGroup(groupId int, memberList []int) error
	ShareGroupLink(groupId int) (string, error)
}
```

- 貼文
```go
type IPost interface {
	NewPost(int, *Post) error
	UpdatePost(int, *Post) error
	SharePostLink(int) (string,error)
}
```
