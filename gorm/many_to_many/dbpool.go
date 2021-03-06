package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Our DemoTable Struct
type DemoTable struct {
	// gorm.Model
	Name  string `gorm:"primary_key"`
	Email string
}

type DBConfig struct {
	User      string
	Password  string
	DBType    string
	DBName    string
	DBAddress string
	DBPort    string
	DBUri     string
}

type OperationDatabase struct {
	DB *gorm.DB
}

type OPDB interface {
	create(name string, email string) error
	queryWithName(name string) (string, error)
	updateEmail(name string, email string) error
	deleteData(name string, email string) error
	Closed() error
	debug()
	transaction() error
}

// An transaction example for `gorm`
func (db *OperationDatabase) transaction() error {
	tx := db.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	// under transaction mode create an object
	if err := tx.Create(&DemoTable{Name: "Jim"}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Create(&DemoTable{Name: "Jim2"}).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error

	return nil
}

func (dbc *DBConfig) NewDBConnection() (OPDB, error) {
	// connection :=
	db, err := gorm.Open(dbc.DBType, dbc.DBUri)
	if err != nil {
		return nil, err
	}
	db = db.AutoMigrate(&DemoTable{})
	return &OperationDatabase{DB: db}, err
}

func NewDBConfiguration(user string, password string, dbtype string, dbname string, dbport string, dbaddress string) *DBConfig {
	return &DBConfig{
		User:      user,
		Password:  password,
		DBType:    dbtype,
		DBName:    dbname,
		DBPort:    dbport,
		DBAddress: dbaddress,
		DBUri:     user + ":" + password + "@tcp(" + dbaddress + ":" + dbport + ")/" + dbname + "?charset=utf8&parseTime=True&loc=Local",
	}
}

func (db *OperationDatabase) Closed() error {
	if err := db.DB.Close(); err != nil {
		return fmt.Errorf("Error happended while closing db : %v\n", err)
	}
	log.Fatalln("Going to close DB")
	return nil
}

// ????????????Debug()?????????????????????SQL??????
func (db *OperationDatabase) debug() {
	db.DB = db.DB.Debug()
}

// ??????CRUD
// Create
func (db *OperationDatabase) create(name string, email string) error {
	var dt = &DemoTable{
		Name:  name,
		Email: email,
	}
	if err := db.DB.Create(dt).Error; err != nil {
		return err
	}
	return nil
}

// Read
func (db *OperationDatabase) queryWithName(name string) (string, error) {
	var dt = &DemoTable{
		Name: name,
	}
	if err := db.DB.Select("email").Find(dt).Error; err != nil {
		return "Can't find the email with " + name, err
	}
	return dt.Email, nil
}

// Update ... ???????????????Read????????????Read??????????????????????????????notes:???gorm?????????????????????????????????updated_at?????????
func (db *OperationDatabase) updateEmail(name string, email string) error {
	return db.DB.First(&DemoTable{Name: name}).Update(&DemoTable{Name: name, Email: email}).Error
}

// Delete ... ??????delete????????????????????????????????????deleteData?????????????????????notes:???gorm???????????????????????????db??????????????????????????????deleted_at?????????
func (db *OperationDatabase) deleteData(name string, email string) error {
	return db.DB.Where("email = ?", email).Delete(&DemoTable{}).Error
}
