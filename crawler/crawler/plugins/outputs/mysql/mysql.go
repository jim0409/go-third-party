package mysql

import (
	"log"
	"time"

	"github.com/jimweng/crawler/crawler/plugins/outputs"
	"github.com/jimweng/crawler/crawler/utils"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type operationDatabase struct {
	DB *gorm.DB
}

type dbOperationInterface interface {
	Write(*[]*utils.PKGContent) error
	debug()
}

type SQLConfig struct {
	DBName        string
	DBPort        string
	DBAddr        string
	User          string
	Password      string
	DBType        string
	MaxIdleConns  int
	MaxOpenConns  int
	KeepAlive     int
	ConnectionUrl string
	DBClient      dbOperationInterface
}

func (opdb *operationDatabase) debug() {
	opdb.DB = opdb.DB.Debug()
}

func (s *SQLConfig) debug() {

}

func (opdb *operationDatabase) Write(points *[]*utils.PKGContent) error {
	// 抓取primary_key; 使用primary_key來判斷是否有建過該record，有的話更新。沒有則創建
	// 要如何用batch 塞資料? https://github.com/jinzhu/gorm/issues/255
	for _, pt := range *points {
		// 1.先create，報錯後再update
		// log.Printf("C!; The value of name:%v\tparent:%v\tsynopsis:%v\thref:%v\n", pt.Name, pt.Parent, pt.Synopsis, pt.Href)
		if err := opdb.DB.Create(pt).Error; err != nil {
			// log.Printf("U!; The value of name:%v\tparent:%v\tsynopsis:%v\thref:%v\n", pt.Name, pt.Parent, pt.Synopsis, pt.Href)
			if err := opdb.DB.First(&utils.PKGContent{Name: pt.Name}).Update(pt).Error; err != nil {
				log.Printf("E!; The value of name:%v\tparent:%v\tsynopsis:%v\thref:%v\n", pt.Name, pt.Parent, pt.Synopsis, pt.Href)
				// return err
			}
		}
		// // 2.先update，報錯後在create
		// if err := opdb.DB.First(&utils.PKGContent{Name: pt.Name}).Update(pt).Error; err != nil {
		// 	if err := opdb.DB.Create(pt).Error; err != nil {
		// 		return fmt.Errorf("Some error happened with msg : %v\n", err)
		// 	}
		// }
	}
	return nil
}

func (s *SQLConfig) newDBConnection() (dbOperationInterface, error) {
	if db, err := gorm.Open(s.DBType, s.ConnectionUrl); err != nil {
		return nil, err
	} else {
		db.AutoMigrate(&utils.PKGContent{})
		// initialize DB settings for connection max nums 10 and keep-alive
		if dbconfig := db.DB(); dbconfig != nil {
			dbconfig.SetMaxOpenConns(s.MaxOpenConns)
			dbconfig.SetMaxIdleConns(s.MaxIdleConns)
			dbconfig.SetConnMaxLifetime(time.Duration(s.KeepAlive))
		}
		return &operationDatabase{DB: db}, nil
	}
}

func (s *SQLConfig) newConnection() {
	connectionUrl := s.User + ":" + s.Password + "@tcp(" + s.DBAddr + ":" + s.DBPort + ")/" + s.DBName + "?charset=utf8&parseTime=True&loc=Local"
	s.ConnectionUrl = connectionUrl
}

func (s *SQLConfig) closeDB(db *gorm.DB) error {
	if err := db.Close(); err != nil {
		return err
	}
	return nil
}

func (s *SQLConfig) Write(points *[]*utils.PKGContent) error {
	if s.ConnectionUrl == "" {
		s.newConnection()
	}
	db, err := s.newDBConnection()
	if err != nil {
		log.Fatalf("Error happened while connecting to DB: %v", err)
		return err
	}

	if err = db.Write(points); err != nil {
		return err
	}

	// if err = db.insertData(points); err != nil {
	// 	return err
	// }

	return nil
}

func init() {
	outputs.Add("mysql", func() utils.Output {
		return &SQLConfig{}
	})
}
