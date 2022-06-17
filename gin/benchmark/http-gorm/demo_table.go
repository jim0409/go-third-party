package main

import (
	"fmt"

	"gorm.io/gorm"
)

type tbMethod interface {
	CRUD(method string, name string, odds int) error
}

type Record struct {
	gorm.Model
	// Name  string `gorm:"primary_key;unique"`
	Integer int
	String  string
}

func (db *Operation) CleanAll() error {
	db.DB.Exec("drop table demo_tables")
	return nil
}

// 實做CRUD
// Create
func (db *Operation) CRUD(method string, str string, integer int) error {

	r := &Record{
		Integer: integer,
		String:  str,
	}

	switch method {
	case "c":
		if err := db.DB.Create(r).Error; err != nil {
			return err
		}
		return nil
	}

	return fmt.Errorf("no such method")
}
