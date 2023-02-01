package main

import (
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func MockInit() OPDB {
	mysqlAddr := "127.0.0.1"
	mysqlPort := "3306"
	mysqlOpDB := "upload_file"
	mysqlUsr := "root"
	mysqUsrPwd := "root"

	newDB := NewDBConfiguration(mysqlUsr, mysqUsrPwd, "mysql", mysqlOpDB, mysqlPort, mysqlAddr)
	db, err := newDB.NewDBConnection()
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func TestInsertOneRecord(t *testing.T) {
	db := MockInit()
	defer db.Closed()
	id, err := db.InsertOneRecord("jim", "file1", "123", 100, 1)
	assert.Nil(t, err)
	assert.Equal(t, id, 1)

	_, err = db.InsertOneRecord("jim", "file1", "123", 100, 1)
	assert.NotNil(t, err)
}

func TestFindUploadDetailByFileName(t *testing.T) {
	db := MockInit()
	defer db.Closed()
	file, err := db.FindUploadDetailByFileName("123", "file1")
	assert.Nil(t, err)
	if file.ID == 0 {
		fmt.Println("empty")
	} else {
		fmt.Println(file)
	}
}

func TestUpdateFileDetails(t *testing.T) {
	db := MockInit()
	defer db.Closed()

	err := db.UpdateFileDetails("123", "file1", "file1-000.png.bak", 1)
	assert.Nil(t, err)
}

func TestFindUploadDetailById(t *testing.T) {
	db := MockInit()
	defer db.Closed()

	ids := []int{1, 2, 3, 4, 5, 6, 7, 8}
	files, err := db.FindUploadDetailById(ids)
	assert.Nil(t, err)
	for _, file := range *files {
		fmt.Println(file)
	}
}

func TestMergeFiles(t *testing.T) {
	filename := "files/auto.mp4"
	chunkfiles := []string{
		"files/jim_split-auto.mp4aa_8_1",
		"files/jim_split-auto.mp4ab_8_2",
		"files/jim_split-auto.mp4ac_8_3",
		"files/jim_split-auto.mp4ad_8_4",
		"files/jim_split-auto.mp4ae_8_5",
		"files/jim_split-auto.mp4af_8_6",
		"files/jim_split-auto.mp4ag_8_7",
		"files/jim_split-auto.mp4ah_8_8",
	}

	err := MergeChunkFiles(filename, chunkfiles)
	assert.Nil(t, err)
}
