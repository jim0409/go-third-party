package main

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func MockInit() OPDB {
	mysqlAddr := "127.0.0.1"
	mysqlPort := "3306"
	mysqlOpDB := "group"
	mysqlUsr := "root"
	mysqUsrPwd := "root"

	newDB := NewDBConfiguration(mysqlUsr, mysqUsrPwd, "mysql", mysqlOpDB, mysqlPort, mysqlAddr)
	db, err := newDB.NewDBConnection()
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func TestNewMember(t *testing.T) {
	db := MockInit()
	defer db.Closed()

	err := db.NewMember("jim1", "password")
	assert.Nil(t, err)
	err = db.NewMember("jim2", "password")
	assert.Nil(t, err)
	err = db.NewMember("jim3", "password")
	assert.Nil(t, err)
}

func TestNewGroup(t *testing.T) {
	db := MockInit()
	defer db.Closed()

	err := db.NewGroup(1, "group1")
	assert.Nil(t, err)
	err = db.NewGroup(2, "group1")
	assert.NotNil(t, err)
	err = db.NewGroup(3, "group1")
	assert.NotNil(t, err)
}

func TestAddMembersToGroup(t *testing.T) {
	db := MockInit()
	defer db.Closed()
	ids := []int{1, 2, 3}
	err := db.AddMembersToGroup(1, ids)
	assert.Nil(t, err)
}

func CreateNewPost(ctx string) *Post {
	return &Post{}
}

func TestAddNewPost(t *testing.T) {
	db := MockInit()
	defer db.Closed()
	usrId := 1
	p := &Post{
		Content: "test content",
	}
	err := db.NewPost(usrId, p)
	assert.Nil(t, err)
}

func TestUpdatePost(t *testing.T) {
	db := MockInit()
	defer db.Closed()
	usrId := 2
	p := &Post{
		ID:      1,
		Content: "update content",
	}
	err := db.UpdatePost(usrId, p)
	assert.Nil(t, err)
}
