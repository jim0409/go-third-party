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

	err := db.UpdateFileDetails("123", "file1", "file1-000.png.bak", 1, 1)
	assert.Nil(t, err)
}

func TestFindUploadDetailByMd5Values(t *testing.T) {
	db := MockInit()
	defer db.Closed()

	md5s := []string{
		"eb02a78c7158e3cfeeeb2989c23d0920",
		"f7a9cd4cf188f4737a17fba0b58268ee",
		"0417f368ad3d98f048d609c6b7961bd5",
		"0394186975fbdaadcce19313a3c368dd",
		"6dcf4aea79fb898599ea0b10064654ba",
		"10ddea23cda77b8d1efda93aabc656cd",
		"f51f84bd33a4a8f6c663a6d4d701f248",
		"f10b0690de37e097054ca28e11be4462",
	}
	chunkfiles, err := db.FindUploadDetailByMd5Values(md5s)
	assert.Nil(t, err)
	for _, chunkfile := range *chunkfiles {
		fmt.Println(chunkfile)
	}
}

func TestAddFileToList(t *testing.T) {
	db := MockInit()
	defer db.Closed()

	filename := "auto.mp4"
	owner := "jim"
	md5s := []string{
		"eb02a78c7158e3cfeeeb2989c23d0920",
		"f7a9cd4cf188f4737a17fba0b58268ee",
		"0417f368ad3d98f048d609c6b7961bd5",
		"0394186975fbdaadcce19313a3c368dd",
		"6dcf4aea79fb898599ea0b10064654ba",
		"10ddea23cda77b8d1efda93aabc656cd",
		"f51f84bd33a4a8f6c663a6d4d701f248",
		"f10b0690de37e097054ca28e11be4462",
	}
	chunkfiles, err := db.FindUploadDetailByMd5Values(md5s)
	assert.Nil(t, err)

	err = db.AddFileToList(filename, owner, chunkfiles)
	assert.Nil(t, err)
}

func TestMergeFiles(t *testing.T) {
	filename := "files/auto.mp4"
	db := MockInit()
	defer db.Closed()

	md5s := []string{
		"eb02a78c7158e3cfeeeb2989c23d0920",
		"f7a9cd4cf188f4737a17fba0b58268ee",
		"0417f368ad3d98f048d609c6b7961bd5",
		"0394186975fbdaadcce19313a3c368dd",
		"6dcf4aea79fb898599ea0b10064654ba",
		"10ddea23cda77b8d1efda93aabc656cd",
		"f51f84bd33a4a8f6c663a6d4d701f248",
		"f10b0690de37e097054ca28e11be4462",
	}
	chunkfiles, err := db.FindUploadDetailByMd5Values(md5s)
	assert.Nil(t, err)

	err = MergeChunkFiles(filename, "jim", chunkfiles)
	assert.Nil(t, err)
}
