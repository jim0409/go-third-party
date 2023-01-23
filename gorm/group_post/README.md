# intro

練習撰寫 FB 社團相關的 api


# 功能介紹
1. 創建會員 .. <IMember>
2. 獲取會員的社團 .. <IMember>

1. 創建社團 .. <IGroup>
2. 修改社團資訊 .. <IGroup>
3. 加入社團 .. <IGroup>
4. 等待加入社團人員 .. <IGroup>
5. 邀請成員加入社團 .. <IGroup>
6. 將成員踢出社團 .. <IGroup>
7. 分享社團連結 .. <IPost>

1. 社團發文 .. <IPost>
2. 獲取該社團的所有發文 .. <IPost>
3. 修改社團發文 .. <IPost>
4. 分享社團發文連結 .. <IPost>
5. 轉發社團發文 .. <IPost>


# 實現功能

- 會員
```go
type IMember interface {
	NewMember(name string, idfy string) error
	GetMemberGroupIndexs(usrId int) (*[]int, error)
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
	NewPost(usrId int, p *Post) error
	GetGroupPost(groupId int) (*[]Post, error)
	UpdatePost(usrId int, p *Post) error
	SharePostLink(postId int) (string, error)
	CopyGroupPost(usrId int, postId int, dstGroupId int) error
}
```
