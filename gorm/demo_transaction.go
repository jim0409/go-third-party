package main

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

type bankImpl interface {
	CreateBankAccount(string, uint64) error  // 創建一個帳號對應 balance 紀錄
	Withdraw(string, uint64) (uint64, error) // 使用者取款
	Deposit(string, uint64) (uint64, error)  // 使用者存款
	Balance(string) (uint64, error)          // 使用者存款
	// Transfer(string, string, uint64) (map[string]uint64, error)
}

type BankUsr struct {
	gorm.Model
	Name    string `gorm:"primary_key;unique"`
	Status  string `gorm:"type:varchar(10);comment:'帳戶狀態';default: active"`
	Balance uint64
}

type Records struct {
	gorm.Model
	history string
}

var tableName = "users"

func (b BankUsr) TableName() string {
	return tableName
}

// An transaction for user withdraw
/*
	創建一個 history records 紀錄
	創建一個銀行帳號，帶入預設金額
*/
func (db *Operation) CreateBankAccount(name string, money uint64) error {
	u := BankUsr{
		Name:    name,
		Balance: money,
	}
	tx := db.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Create(&u).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	// 直接回傳創建好的 id，不用再做 last_id 的查詢 (`SELECT LAST_INSERT_ID()`)
	// https://stackoverflow.com/a/28776926
	log.Println(u.ID)

	return nil
}

func (db *Operation) Deposit(name string, money uint64) (uint64, error) {
	tx := db.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return 0, err
	}

	if err := tx.Table(tableName).Updates(&BankUsr{Name: name, Balance: money}).Error; err != nil {
		tx.Rollback()
		return 0, err
	}

	if err := tx.Commit().Error; err != nil {
		return 0, err
	}

	return uint64(money), nil
}

func (db *Operation) Balance(name string) (uint64, error) {
	var money uint64
	tx := db.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return 0, err
	}

	if err := tx.Table(tableName).Select("balance").Where("name = ? and deleted_at IS NULL", name).Scan(&money).Error; err != nil {
		tx.Rollback()
		return 0, err
	}

	if err := tx.Commit().Error; err != nil {
		return 0, err
	}

	return money, nil
}

func (db *Operation) Withdraw(name string, money uint64) (uint64, error) {
	var wealth uint64
	tx := db.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return 0, err
	}

	err := tx.Debug().Table(tableName).Select("balance").Where("name = ? and balance > ? and deleted_at IS NULL", name, money).Scan(&wealth).Error
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	if wealth <= 0 {
		return 0, fmt.Errorf("balance is not enough")
	}

	err = tx.Debug().Table(tableName).Updates(&BankUsr{Name: name, Balance: wealth - money}).Error
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	if err := tx.Commit().Error; err != nil {
		return 0, err
	}

	return wealth - money, nil
}
