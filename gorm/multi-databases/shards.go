package main

import "gorm.io/gorm"

type MessageTable struct {
	gorm.Model
	Name string `gorm:"column:name"`
	Age  string `gorm:"column:age"`

	TableName string `gorm:"-"`
}

/*
1. 創建訊息群組
2. 增加群組人員訊息
3. 管制DB查詢位置
*/

type GroupMessage interface {
	createGroupMsgTabel(string) error
	InsertRecrods(string, string, string) error
}

func (o *Operation) createGroupMsgTabel(tbname string) error {
	return o.DB.Table(tbname).AutoMigrate(&MessageTable{TableName: tbname})
}

func (o *Operation) InsertRecrods(tbname, name string, age string) error {
	return o.DB.Table(tbname).Create(&MessageTable{Name: name, Age: age}).Error
}
