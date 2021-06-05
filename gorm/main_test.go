package main

import (
	"log"
	"os"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func init() {
	clearTestEnv()
}

func clearTestEnv() {
	log.Println("Clear test env")
	if _, err := os.Stat("/tmp/gorm.db"); !os.IsNotExist(err) {
		log.Println("Remove Origin gorm.db Files")
		os.Remove("/tmp/gorm.db")
	}
}

// Set mock DB env
func mockTestEnv() (OPDB, error) {
	var dbc = DBConfig{}
	dbc.DBUri = "/tmp/gorm.db"
	dbc.DBType = "sqlite3"
	opdb, err := dbc.NewDBConnection()
	return opdb, err
}

func TestNewConnection(t *testing.T) {
	var (
		user     = "jim"
		password = "pw"
		dbtype   = "sqlite"
		dbname   = "demo_db"
		dbport   = "3306"
		dbaddr   = "127.0.0.1"
	)

	NewDBConfig := NewDBConfiguration(user, password, dbtype, dbname, dbport, dbaddr)
	assert.Equal(t, user, NewDBConfig.User)
	assert.Equal(t, password, NewDBConfig.Password)
	assert.Equal(t, dbtype, NewDBConfig.DBType)
	assert.Equal(t, dbname, NewDBConfig.DBName)
	assert.Equal(t, dbport, NewDBConfig.DBPort)
	assert.Equal(t, dbaddr, NewDBConfig.DBAddress)
	assert.Equal(t, "jim:pw@tcp(127.0.0.1:3306)/demo_db?charset=utf8&parseTime=True&loc=Local", NewDBConfig.DBUri)

}

func TestNewDBOperation(t *testing.T) {
	opdb, err := mockTestEnv()
	assert.Nil(t, err)

	dt := &DemoTable{
		Name:  "jim",
		Email: "email@example.com",
	}
	update_email := "update_email@example.com"

	opdb.debug()

	err = opdb.create(dt.Name, dt.Email)
	assert.Nil(t, err)

	resStr, err := opdb.queryWithName(dt.Name)
	assert.Equal(t, dt.Email, resStr)
	assert.Nil(t, err)

	err = opdb.updateEmail(dt.Name, update_email)
	assert.Nil(t, err)

	resStr, err = opdb.queryWithName(dt.Name)
	assert.Equal(t, update_email, resStr)

	err = opdb.deleteData(dt.Name, update_email)
	assert.Nil(t, err)

	resStr, err = opdb.queryWithName(dt.Name)
	assert.Equal(t, "Can't find the email with jim", resStr)
	assert.Error(t, err)

}
