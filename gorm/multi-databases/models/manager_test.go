package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func init() {
	var err error

	dbc := NewDBConfiguration("jim", "password", "mysql", "db", "3306", "127.0.0.1")
	db, err = gorm.Open(mysql.Open(dbc.DBUri), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if err := db.Migrator().DropTable(&NodeInfo{}, &GroupInNodes{}); err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(&NodeInfo{}); err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(&GroupInNodes{}); err != nil {
		panic(err)
	}
}

func TestMain(t *testing.T) {
	testAddNodeInfos(t)
	testMockGroupInNodes(t)
	testQueryGroupLoc(t)
	testNodeStatic(t)
}

func testAddNodeInfos(t *testing.T) {
	mdb := &MainDb{
		SettingDB: &Operation{DB: db},
	}

	assert.Nil(t, mdb.SettingDB.AddNodeInfos([]NodeInfo{
		NodeInfo{
			User:     "jim",
			Password: "password",
			Type:     "mysql",
			Database: "message",
			Port:     "3301",
			Address:  "127.0.0.1",
		},
		NodeInfo{
			User:     "jim",
			Password: "password",
			Type:     "mysql",
			Database: "message",
			Port:     "3302",
			Address:  "127.0.0.1",
		},
		NodeInfo{
			User:     "jim",
			Password: "password",
			Type:     "mysql",
			Database: "message",
			Port:     "3303",
			Address:  "127.0.0.1",
		},
	}))
}

func testMockGroupInNodes(t *testing.T) {
	mdb := &MainDb{
		SettingDB: &Operation{DB: db},
	}

	assert.Nil(t, mdb.SettingDB.AddGroupInNodes("msg-0", 1))
	assert.Nil(t, mdb.SettingDB.AddGroupInNodes("msg-1", 2))
	assert.Nil(t, mdb.SettingDB.AddGroupInNodes("msg-2", 3))
	assert.Nil(t, mdb.SettingDB.AddGroupInNodes("msg-3", 3))
	assert.Nil(t, mdb.SettingDB.AddGroupInNodes("msg-4", 3))
}

func testQueryGroupLoc(t *testing.T) {
	mdb := &MainDb{
		SettingDB: &Operation{DB: db},
	}

	loc, err := mdb.SettingDB.QueryGroupLoc("msg-0")
	assert.Nil(t, err)
	assert.Equal(t, 1, loc)

	loc, err = mdb.SettingDB.QueryGroupLoc("msg-1")
	assert.Nil(t, err)
	assert.Equal(t, 2, loc)

	loc, err = mdb.SettingDB.QueryGroupLoc("msg-2")
	assert.Nil(t, err)
	assert.Equal(t, 3, loc)

	loc, err = mdb.SettingDB.QueryGroupLoc("msg-999")
	assert.Nil(t, err)
	assert.Equal(t, -1, loc)
}

func testNodeStatic(t *testing.T) {
	mdb := &MainDb{
		SettingDB: &Operation{DB: db},
	}

	loc, err := mdb.SettingDB.NodeStatic()
	if err != nil {
		panic(err)
	}
	assert.Equal(t, loc, 3)
}

func TestCreateMessage(t *testing.T) {

}

func TestReadMessage(t *testing.T) {

}
