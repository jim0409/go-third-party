package main

import (
	"time"

	"gorm.io/gorm"
)

type FileUploadDetail struct {
	ID            int    `gorm:"primaryKey;autoIncrement;"`
	Username      string `gorm:"type:varchar(32);comment:使用者名稱"`
	FileName      string `gorm:"index;type:varchar(64);comment:檔案名稱"`
	Size          int64  `gorm:"type:int(64);comment:檔案大小"`
	ChunkFilename string `gorm:"unique;type:varchar(64);伺服器端分片的檔案名稱"`
	Md5           string `gorm:"unique;index;type:varchar(255);md5值"`
	IsUploaded    int    `gorm:"type:tinyint(8);default:0;comment:0,還沒上傳 1,已上傳"`
	ChunkNum      int    `gorm:"type:tinyint(8);default:0;comment:分片序號"`
	TotalChunks   int    `gorm:"type:tinyint(8);default:0;comment:總分片數"`
	UidFile       string `gorm:"type:varchar(1024);comment:定義上傳的檔案唯一識別名稱"`
	Url           string `gorm:"type:varchar(255);comment:上傳網址"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
}

type IFileUploadDetail interface {
	InsertOneRecord(usrname string, filename string, md5 string, size int64, totalchunks int) (int, error)
	FindUploadDetailByFileName(md5 string, filename string) (*FileUploadDetail, error)
	UpdateFileDetails(md5 string, filename string, chunkfilename string, chunknum int) error
}

// InsertOneRecord: 保存一條紀錄
func (db *Operation) InsertOneRecord(usrname string, filename string, md5 string, size int64, totalchunks int) (int, error) {
	file := &FileUploadDetail{
		Username:    usrname,
		FileName:    filename,
		Size:        size,
		Md5:         md5,
		TotalChunks: totalchunks,
	}

	return file.ID, db.DB.Table("file_upload_details").Create(file).Error
}

// FindUploadDetailByFileName: 根據 md5, filename 查詢文件
func (db *Operation) FindUploadDetailByFileName(md5 string, filename string) (*FileUploadDetail, error) {
	file := &FileUploadDetail{}
	err := db.DB.Table("file_upload_details").
		Select("*").
		Where("md5 = ? AND file_name = ?", md5, filename).Scan(file).Error
	if err != nil {
		return nil, err
	}

	return file, nil
}

// UpdateFileDetails: 更新上傳檔案細節
func (db *Operation) UpdateFileDetails(md5 string, filename string, chunkfilename string, chunknum int) error {
	file, err := db.FindUploadDetailByFileName(md5, filename)
	if err != nil {
		return err
	}

	file.IsUploaded = 1
	file.ChunkNum = chunknum // 定義上傳的檔案唯一識別名稱
	file.ChunkFilename = chunkfilename
	file.UidFile = "todo - need to specified the upload file name"

	return db.DB.Table("file_upload_details").Updates(file).Where("md5 = ? AND file_name = ?", file.Md5, file.FileName).Error
}
