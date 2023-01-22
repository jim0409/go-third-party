package main

import (
	"time"

	"gorm.io/gorm"
)

// 紀錄登入者信息
type Member struct {
	ID         int     `gorm:"primaryKey;autoIncrement;"`
	Nickname   string  `gorm:"unique;type:varchar(32);comment:使用者名稱"`
	Identify   string  `gorm:"type:varchar(32);comment:bcrypt對照"`
	GroupInfos []Group `gorm:"many2many:group_members;"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt
}

// 紀錄社團資訊
type Group struct {
	ID           int      `gorm:"primaryKey;autoIncrement;"`
	Name         string   `gorm:"unique;type:varchar(32);comment:群組名稱"`
	Owner        string   `gorm:"type:varchar(32);comment:群組創建者"`
	Members      []Member `gorm:"many2many:group_members;"`
	AwaitMembers []Member `gorm:"many2many:group_members;"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
}

// 紀錄社團內每個人的身分
type GroupMember struct {
	GroupID  int    `gorm:"primaryKey"`
	MemberID int    `gorm:"primaryKey"`
	Role     string `gorm:"type:tinyint(8);comment:腳色"`
}

// 紀錄貼文
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

// 通用方法 //

// GetMemberViaNickname: 透過 nickname 獲取使用者資訊
func (db *Operation) GetMemberViaNickname(name string) (*Member, error) {
	m := &Member{}
	err := db.DB.Table("members").Select("*").Where("nickname = ?", name).Scan(m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

// GetMemberViaId: 透過 nickname 獲取使用者資訊
func (db *Operation) GetMemberViaId(id int) (*Member, error) {
	m := &Member{}
	err := db.DB.Table("members").Select("*").Where("id = ?", id).Scan(m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

// GetMemberViaListIds: 透過 id list 獲取使用者資訊
func (db *Operation) GetMemberViaListIds(list []int) (*[]Member, error) {
	members := []Member{}
	err := db.DB.Table("members").Select("*").Where("id IN ?", list).Scan(&members).Error
	if err != nil {
		return nil, err
	}
	return &members, nil
}

// 定義功能
// 會員的 Interface
type IMember interface {
	NewMember(name string, idfy string) error
}

// NewMember: 增加新成員; idfy: bcrypt endcoding
func (db *Operation) NewMember(name string, idfy string) error {
	m := &Member{
		Nickname:   name,
		Identify:   idfy,
		GroupInfos: nil,
	}
	return db.DB.Table("members").Create(m).Error
}

// 社團的 Interface
type IGroup interface {
	NewGroup(userId int, name string) error
	UpdateGroup(usrId int, gp *Group) error
	ApplyForGroup(userId int, groupId int) error
	AwaitJoinGroupList(groupId int) ([]int, error)
	AddMembersToGroup(groupId int, memberList []int) error
	KickMemberOutOfGroup(groupId int, memberList []int) error
	ShareGroupLink(groupId int) (string, error)
}

// NewGroup: 創建一個新的 group
func (db *Operation) NewGroup(userId int, name string) error {
	// 1. check user id, if member not existed return err
	m, err := db.GetMemberViaId(userId)
	if err != nil {
		return err
	}

	// 2. assign member into initial value and create new group
	gp := &Group{
		Owner:   m.Nickname,
		Name:    name,
		Members: []Member{*m},
	}
	return db.DB.Table("groups").Create(gp).Error
}

func (db *Operation) UpdateGroup(usrId int, gp *Group) error {
	return nil
}

func (db *Operation) ApplyForGroup(userId int, groupId int) error {
	return nil
}

func (db *Operation) AwaitJoinGroupList(groupId int) ([]int, error) {
	return nil, nil
}

// AddMembersToGroup: 增加 group 成員
func (db *Operation) AddMembersToGroup(groupId int, list []int) error {
	// 1. retrieve all the members in list
	members, err := db.GetMemberViaListIds(list)
	if err != nil {
		return err
	}
	// 2. updates group_members via groups index
	gp := &Group{
		ID:      groupId,
		Members: *members,
	}

	return db.DB.Table("groups").Updates(gp).Error
}

func (db *Operation) KickMemberOutOfGroup(groupId int, memberList []int) error {
	return nil
}

func (db *Operation) ShareGroupLink(groupId int) (string, error) {
	return "", nil
}

// 貼文的 Interface
type IPost interface {
	NewPost(usrId int, p *Post) error
	UpdatePost(usrId int, p *Post) error
	SharePostLink(postId int) (string, error)
}

// NewPost: 增加一篇貼文; 需要再使用前，先檢查 post 的內容以及對應的
func (db *Operation) NewPost(usrId int, p *Post) error {
	m, err := db.GetMemberViaId(usrId)
	if err != nil {
		return err
	}
	p.Owner = m.Nickname

	return db.DB.Table("posts").Create(p).Error
}

// UpdatePost: 更新貼文; 不改變 sharelink, 不改 like/ disklike 數
func (db *Operation) UpdatePost(usrId int, p *Post) error {
	// 1. check whether usr existed in here
	m, err := db.GetMemberViaId(usrId)
	if err != nil {
		return err
	}
	p.Editor = m.Nickname

	return db.DB.Table("posts").Updates(p).Where("id = ?", p.ID).Error
}

func (db *Operation) SharePostLink(postId int) (string, error) {
	return "", nil
}
