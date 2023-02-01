package main

import (
	"time"

	"gorm.io/gorm"
)

type FileList struct {
	ID        int    `gorm:"primaryKey;autoIncrement;"`
	Owner     string `gorm:"type:varchar(32);comment:使用者名稱"`
	Size      int64  `gorm:"type:int(64);comment:檔案大小"`
	Url       string `gorm:"type:varchar(255);comment:下載網址"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type IFileList interface {
	MergeFile(filename string, username string) error
}

func (db *Operation) MergeFile(filelist []int) {
	// db.FindUploadDetailById()

}
