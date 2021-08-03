package main

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
)

type tbMethod interface {
	create(name string, email string) error
	queryWithName(name string) (string, error)
	updateEmail(name string, email string) error
	deleteData(name string, email string) error
	// transaction(string, string) error
	cleanAll() error
}

// Our DemoTable Struct
type DemoTable struct {
	gorm.Model
	// Name  string `gorm:"primary_key;unique"`
	Name  string
	Email string
}

func (db *Operation) cleanAll() error {
	db.DB.Exec("drop table demo_tables")
	return nil
}

// 實做CRUD
// Create
func (db *Operation) create(name string, email string) error {
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
func (db *Operation) queryWithName(name string) (string, error) {
	var dt = &DemoTable{
		Name: name,
	}
	if err := db.DB.Select("email").Find(dt).Error; err != nil {
		return "Can't find the email with " + name, err
	}
	return dt.Email, nil
}

// Update ... 更新相當於Read以後在把Read的資料改成新的資料；notes:在gorm裡面，更新以後也會更新updated_at的時間
func (db *Operation) updateEmail(name string, email string) error {
	// return db.DB.First(&DemoTable{Name: name}).Updates(&DemoTable{Name: name, Email: email}).Error
	return db.DB.Updates(&DemoTable{Name: name, Email: email}).Where("name = ? and deleted_at is NULL", name).Error
}

// Delete ... 因為delete已經有預設方法，這邊改用deleteData來宣告該函數；notes:在gorm裡面刪除不是代表從db完全移除。而是去更改deleted_at的時間
func (db *Operation) deleteData(name string, email string) error {
	// log.Printf("The %s's Email has been delete (%s)", name, db.DB.Delete(&DemoTable{Name: name, Email: email}).Value)
	if err := db.DB.Where("email = ?", email).Delete(&DemoTable{}).Error; err != nil {
		return fmt.Errorf("Encount Error with no data to delete %v\n", err)
	}
	return nil
}

// 給定一連串的條件做 注入 or 查詢
// https://zhiruchen.github.io/2017/08/31/bulk-insert-bulk-query-with-gorm/
/*
[
	{"name": "jim1", "email": "example.com"},
	{"name": "jim2", "email": "example.com"},
	{"name": "jim2", "email": "example.com"}
]
*/
func (db *Operation) bulkInsert(recrods []map[string]interface{}) error {
	valueStrings := []string{}
	valueArgs := []interface{}{}

	for _, f := range recrods {
		valueStrings = append(valueStrings, "(?, ?, ?)")

		valueArgs = append(valueArgs, f["name"])
		valueArgs = append(valueArgs, f["email"])
	}

	smt := `INSERT INTO record_t(id, field1, field2) VALUES %s ON DUPLICATE KEY UPDATE field2=VALUES(field2)`

	smt = fmt.Sprintf(smt, strings.Join(valueStrings, ","))

	tx := db.DB.Begin()
	if err := tx.Exec(smt, valueArgs...).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

/*
[
	{"name": "jim1"},
	{"name": "jim2"}
]
*/
func (db *Operation) bulkQuery(filters []map[string]interface{}) (interface{}, error) {
	rs := []*DemoTable{}

	placeHolders := []string{}
	args := []interface{}{}

	for _, filter := range filters {
		placeHolders = append(placeHolders, "?")
		args = append(args, filter["name"])
	}

	sql := `select name, email from demo_table where name in(%s)`
	sql = fmt.Sprintf(sql, strings.Join(placeHolders, ","))

	if err := db.DB.Begin().Raw(sql, args...).Scan(&rs).Error; err != nil {
		return nil, err
	}

	return &rs, nil
}
