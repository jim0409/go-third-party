package models

import (
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitConnection() *DBConfig {
	// return NewDBConfiguration("jim", "password", "mysql", "db", "3306", "127.0.0.1")
	return NewDBConfiguration("jim", "password", "mysql", "db", "3306", "10.200.6.99")
}
func TestDropData(t *testing.T) {
	dbc := InitConnection()
	if err := DropTable(dbc); err != nil {
		panic(err)
	}
}

func DropTable(dbc *DBConfig) error {
	db, err := gorm.Open(mysql.Open(dbc.DBUri), &gorm.Config{})
	if err != nil {
		return err
	}
	return db.Migrator().DropTable(&NodeInfo{}, &GroupInDB{})
}

func TestMockData(t *testing.T) {
	dbc := InitConnection()
	if err := MockNodesInfo(dbc); err != nil {
		panic(err)
	}
}

func MockNodesInfo(dbc *DBConfig) error {
	db, err := gorm.Open(mysql.Open(dbc.DBUri), &gorm.Config{})
	if err != nil {
		return err
	}

	if err := db.AutoMigrate(&NodeInfo{}, &GroupInDB{}); err != nil {
		return err
	}

	dbs := []NodeInfo{
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
	}

	return db.Create(dbs).Error
}
