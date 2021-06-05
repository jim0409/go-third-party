package main

import (
	"fmt"
	"log"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/gorm"
)

// Our DemoTable Struct
type DemoTable struct {
	gorm.Model
	Name  string `gorm:"primary_key"`
	Email string
}

func (db *OperationDatabase) cleanAll() error {
	db.DB.Exec("drop table demo_tables")
	return nil
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

func (db *OperationDatabase) Closed() error {
	if err := db.DB.Close(); err != nil {
		return fmt.Errorf("Error happended while closing db : %v\n", err)
	}
	log.Fatalln("Going to close DB")
	return nil
}

// 透過使用Debug()可以轉譯語言為SQL語法
func (db *OperationDatabase) debug() {
	db.DB = db.DB.Debug()
}

// 實做CRUD
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

// Update ... 更新相當於Read以後在把Read的資料改成新的資料；notes:在gorm裡面，更新以後也會更新updated_at的時間
func (db *OperationDatabase) updateEmail(name string, email string) error {
	// log.Printf("The %s's Email has been update to %s", name, db.DB.First(&DemoTable{Name: name}).Update(&DemoTable{Name: name, Email: email}).Value)
	if err := db.DB.First(&DemoTable{Name: name}).Update(&DemoTable{Name: name, Email: email}).Error; err != nil {
		return err
	}
	return nil
}

// Delete ... 因為delete已經有預設方法，這邊改用deleteData來宣告該函數；notes:在gorm裡面刪除不是代表從db完全移除。而是去更改deleted_at的時間
func (db *OperationDatabase) deleteData(name string, email string) error {
	// log.Printf("The %s's Email has been delete (%s)", name, db.DB.Delete(&DemoTable{Name: name, Email: email}).Value)
	if err := db.DB.Where("email = ?", email).Delete(&DemoTable{}).Error; err != nil {
		log.Fatal("Encount Error with no data to delete")
		return err
	}
	return nil
}
