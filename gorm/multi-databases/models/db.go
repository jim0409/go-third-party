package models

import (
	"fmt"

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

// --------- for manager db implement methods
type MainDB interface {
	Closed() error
	RetriveNodes() ([]NodeInfo, error)
	AddNodeInfos([]NodeInfo) error
	AddGroupInNodes(string, int) error
	QueryGroupLoc(string) (int, error)
	NodeStatic() (int, error)

	migrate(...interface{}) error
}

func (c *DBConfig) NewMainDBConnection() (MainDB, error) {
	db, err := gorm.Open(mysql.Open(c.DBUri), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	d, err := db.DB()
	if err != nil {
		return nil, err
	}

	// TODO: 拉出來，放到設定檔做設定
	d.SetMaxOpenConns(200)
	d.SetMaxIdleConns(50)

	return &Operation{DB: db}, err
}

// ----------- for shards db implement methods
type OPDB interface {
	Closed() error
	Debug()

	migrate(...interface{}) error

	GroupMessage
}

func (c *DBConfig) NewDBConnection() (OPDB, error) {
	db, err := gorm.Open(mysql.Open(c.DBUri), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	d, err := db.DB()
	if err != nil {
		return nil, err
	}

	// TODO: 拉出來，放到設定檔做設定
	d.SetMaxOpenConns(200)
	d.SetMaxIdleConns(50)

	return &Operation{DB: db}, err
}

func (o *Operation) migrate(tb ...interface{}) error {
	return o.DB.AutoMigrate(tb...)
}

func (o *Operation) Closed() error {
	db, err := o.DB.DB()
	if err != nil {
		return err
	}
	return db.Close()
}

// 透過使用Debug()可以轉譯語言為SQL語法
func (o *Operation) Debug() {
	o.DB = o.DB.Debug()
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
