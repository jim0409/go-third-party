package main

import "gorm.io/gorm"

type bankImpl interface {
	CreateBankAccount(string, uint64) error // 創建一個帳號對應 balance 紀錄
	// Withdraw(string, uint64) (uint64, error) // 使用者取款
	// Deposit(string, uint64) (uint64, error)  // 使用者存款
	// Balance(string) (uint64, error)          // 使用者存款
	// Transfer(string, string, uint64) (map[string]uint64, error)
}

type BankUsr struct {
	gorm.Model
	Name    string `gorm:"primary_key;unique"`
	Status  string
	Balance uint64
}

type Records struct {
	gorm.Model
	history string
}

// An transaction for user withdraw
/*
	創建一個 history records 紀錄
	創建一個銀行帳號，帶入預設金額
*/
func (db *Operation) CreateBankAccount(name string, money uint64) error {
	tx := db.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	// 開戶 為 1 or 0
	if err := tx.Create(&BankUsr{Name: name, Balance: money}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
