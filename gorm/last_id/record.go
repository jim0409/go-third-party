package main

type UsrIface interface {
	InsertUsr(*Usr) error
}

type Usr struct {
	ID   int
	Name string
}

func (u *Usr) TableName() string {
	return "user"
}

func (o *Operation) InsertUsr(u *Usr) error {
	return o.DB.Table("user").Create(u).Error
}
