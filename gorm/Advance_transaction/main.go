package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// 寫法直接，程序流程明確，但是交易與程序流程嵌入太深，容易遺漏
func FirstTypeOFTransactionOperaion(db *gorm.DB) error {
	var err error
	tx := db.Begin()
	if err := tx.Error; err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // re-throw panic after Rollback
		}
	}()

	if err = tx.Exec("With RAW SQL, do something need to use transaction").Error; err != nil {
		tx.Rollback()
		return err
	}

	if err = tx.Exec("With RAW SQL, do another things need to use transaction").Error; err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit().Error
	return err

}

// 把交易行為從函數中抽出來，不容易溢漏，但是容易造成整段程序流程不清晰
func SecondTypeOfTransactionOperation(db *gorm.DB) error {
	var err error
	tx := db.Begin()
	if err = tx.Error; err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // re-throw panic after Rollback
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit().Error
		}
	}()

	if err = tx.Exec("With RAW SQL, do something need to use transaction").Error; err != nil {
		return err
	}

	if err = tx.Exec("With RAW SQL, do another things need to use transaction").Error; err != nil {
		return err
	}

	return err
}

// 將二的寫法進一步封裝，讓程式可讀性高一點，但是依舊不是很直覺
func ThirdTypeOfTransactionOperation(db *gorm.DB, txFunc func(*gorm.DB) error) error {
	var err error
	tx := db.Begin()
	if err = tx.Error; err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // <- re-throw panic after Rollback
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit().Error
		}
	}()

	err = txFunc(tx)

	return err
}

// 第三種寫法的封裝交易函數
func DoSomethingNeedTransaction(db *gorm.DB) error {
	return ThirdTypeOfTransactionOperation(db, func(tx *gorm.DB) error {
		if err := tx.Exec("With RAW SQL, do something need to use transaction").Error; err != nil {
			return err
		}
		if err := tx.Exec("With RAW SQL, do another things need to use transaction").Error; err != nil {
			return err
		}
		return nil
	})
}

func main() {
	var err error
	fmt.Println(err)
}

// 參考文章的作者寫法
/*
defer tx.Rollback()使得交易回滾。始終必須要執行
當tx.Commit()執行後，tx.Rollback()起到關閉交易的作用
當程序因為某個錯誤終止，tx.Rollback()起到回滾交易，同時關閉交易作用
*/

// 一般場景
func RecommandTypeOfTransactionCommon(db *gorm.DB) error {
	var err error
	tx := db.Begin()
	if err = tx.Error; err != nil {
		return err
	}

	defer tx.Rollback()

	if err = tx.Exec("...").Error; err != nil {
		return err
	}

	err = tx.Commit().Error

	return err
}

// 循環場景: 小交易，每次循環提交一次。在循環內部使的這種寫法，defer不能使用，所以要把交易分離到獨立的函數
func RecommandTypeOfTransactionSmallAndLoop(db *gorm.DB) error {
	var err error
	tx := db.Begin()
	if err = tx.Error; err != nil {
		return err
	}
	defer tx.Rollback()

	for {
		if err = RecommandTypeOfTransactionCommon(tx); err != nil {
			return err
		}
	}

	if err = tx.Commit().Error; err != nil {
		return err
	}

	return err
}
