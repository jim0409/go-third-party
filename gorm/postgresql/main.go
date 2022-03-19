package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Record struct {
	gorm.Model
	// Name  string `gorm:"primary_key;unique"`
	Name  string
	Email string
}

func main() {
	dsn := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("ope db err %v", err))
	}

	if err := db.AutoMigrate(Record{}); err != nil {
		panic(fmt.Sprintf("migrate err %v", err))
	}

	if err := db.Create(&Record{
		Name:  "jim",
		Email: "jim@mail.com",
	}).Error; err != nil {
		panic(fmt.Sprintf("insert record failed %v", err))
	}

}
