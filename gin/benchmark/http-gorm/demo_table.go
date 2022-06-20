package main

import (
	"gorm.io/gorm"
)

type tbMethod interface {
	Create(name string, odds int) error
	Update(name string, odds int) error
}

type Record struct {
	gorm.Model
	// Name  string `gorm:"primary_key;unique"`
	Integer int
	String  string
}

const tblName = "records"

func (r *Record) TableName() string {
	return tblName
}

func (db *Operation) CleanAll() error {
	db.DB.Exec("drop table demo_tables")
	return nil
}

// 實做CRUD
// Create
func (db *Operation) Create(str string, integer int) error {
	r := &Record{
		Integer: integer,
		String:  str,
	}

	if err := db.DB.Create(r).Error; err != nil {
		return err
	}

	return nil
}

// Update
func (db *Operation) Update(str string, integer int) error {
	if err := db.DB.Table(tblName).Where("string = ?", str).Updates(&Record{String: str, Integer: integer}).Error; err != nil {
		return err
	}

	return nil
}
