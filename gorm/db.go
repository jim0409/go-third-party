package main

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBConfig struct {
	User      string
	Password  string
	DBType    string
	DBName    string
	DBAddress string
	DBPort    string
	DBUri     string
}

type Operation struct {
	DB *gorm.DB
}

type OPDB interface {
	Closed() error
	Debug()
	tbMethod
	bankImpl
}

func (dbc *DBConfig) NewDBConnection() (OPDB, error) {
	db, err := gorm.Open(mysql.Open(dbc.DBUri), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if os.Getenv("DEBUG") != "false" {
		err = db.Migrator().DropTable(
			&DemoTable{},
			&BankUsr{},
		)
		if err != nil {
			return nil, err
		}
	}

	d, err := db.DB()
	if err != nil {
		return nil, err
	}

	// 設定連線池，預設最大連線 100 條，閒置 50 條
	d.SetMaxOpenConns(50)
	d.SetMaxIdleConns(50)

	err = db.AutoMigrate(
		&DemoTable{},
		&BankUsr{},
	)
	if err != nil {
		return nil, err
	}

	return &Operation{DB: db}, err
}

func (db *Operation) Closed() error {
	d, err := db.DB.DB()
	if err != nil {
		return err
	}
	return d.Close()
}

// 透過使用Debug()可以轉譯語言為SQL語法
func (db *Operation) Debug() {
	db.DB = db.DB.Debug()
}

func NewDBConfiguration(user, password string, dbtype string, dbname, dbport string, address string) *DBConfig {
	linkUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True", // 會將UTC-time轉成當地時間...自動加8小時
		user, password, address, dbport, dbname,
	)

	return &DBConfig{
		User:      user,
		Password:  password,
		DBType:    dbtype,
		DBName:    dbname,
		DBPort:    dbport,
		DBAddress: address,
		DBUri:     linkUrl,
	}
}
