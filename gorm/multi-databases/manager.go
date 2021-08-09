package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

// TODO: 要處理 main db 的 singleton
/*
var (
	mDB             DBManager
	mainDBTableName = "main"
)
*/

// 紀錄 Group 在哪一個 Shard DB 上
type GroupInDB struct {
	gorm.Model
	GroupID string
	ShardID string
}

// 紀錄 Shards DB 的資訊
type ShardInfo struct {
	gorm.Model
	User     string
	Password string
	Type     string
	Database string
	Port     string
	Address  string
}

type MainDb struct {
	DB      ManagerDB      // 主配置，紀錄 shard dbs 的配置，及動態增加 shards
	ShardDb []OPDB         // 紀錄 shard dbs 的實例
	TabeLoc map[string]int // 反查回 shard db 的位置
}

func (db *Operation) RetriveShards() ([]ShardInfo, error) {
	shards := []ShardInfo{}
	if err := db.DB.Table("shard_infos").Select("*").Scan(&shards).Error; err != nil {
		return nil, err
	}
	return shards, nil
}

func InitMainDB(usr, pwd, tpe, name, port, addr string) DBManager {
	mamindb := NewDBConfiguration(usr, pwd, tpe, name, port, addr)
	mdb, err := mamindb.NewDBMainConnection()
	if err != nil {
		panic(err)
	}
	err = mdb.migrate(&GroupInDB{}, &ShardInfo{})
	if err != nil {
		panic(err)
	}

	return &MainDb{
		DB:      mdb,
		ShardDb: make([]OPDB, 0),
		TabeLoc: make(map[string]int),
	}
}

func (db *Operation) MockShards() {
	dbs := []ShardInfo{
		ShardInfo{
			User:     "jim",
			Password: "password",
			Type:     "mysql",
			Database: "message",
			Port:     "3301",
			Address:  "127.0.0.1",
		},
		ShardInfo{
			User:     "jim",
			Password: "password",
			Type:     "mysql",
			Database: "message",
			Port:     "3302",
			Address:  "127.0.0.1",
		},
		ShardInfo{
			User:     "jim",
			Password: "password",
			Type:     "mysql",
			Database: "message",
			Port:     "3303",
			Address:  "127.0.0.1",
		},
	}

	if err := db.DB.Create(dbs).Error; err != nil {
		fmt.Println(err)
	}
}

/*
TODO:
	改成 query main db 下的 ShardInfo 表
	考慮是否要定期做排查? 以及增加程序內部緩存?
*/
func (m *MainDb) retriveDBs() ([]ShardInfo, error) {
	if os.Getenv("mock") == "true" {
		m.DB.MockShards()
	}
	return m.DB.RetriveShards()
}

func (m *MainDb) StartShardDbs() {
	var dbs []OPDB
	cfgs, err := m.retriveDBs()
	if err != nil {
		panic(err)
	}

	for _, cfg := range cfgs {
		obj := NewDBConfiguration(cfg.User, cfg.Password, cfg.Type, cfg.Database, cfg.Port, cfg.Address)
		db, err := obj.NewDBConnection()
		if err != nil {
			fmt.Printf("err in starup db .. %v\n", err)
		}
		dbs = append(dbs, db)
	}

	m.ShardDb = dbs
}

type DBManager interface {
	StartShardDbs()

	DBWriter
	DBReader
}

type DBWriter interface {
	CreateMessage(string, string, string) error
}

type DBReader interface {
	ReadMessage(string, string, string) error
}

// TODO: 製作決策器 .. 要下決策決定他是去哪一個 shard
func dbDecision(group string) (int, error) {
	id := strings.Split(group, "-")[1]
	// return group_id
	return strconv.Atoi(id)
}

func (o *MainDb) dbSelect(group string) OPDB {
	if t, ok := o.TabeLoc[group]; ok {
		return o.ShardDb[t]
	}

	loc, err := dbDecision(group)
	if err != nil {
		panic(err)
	}

	o.TabeLoc[group] = loc
	if err := o.ShardDb[loc].createGroupMsgTabel(group); err != nil {
		panic(err)
	}

	return o.ShardDb[loc]
}

func (o *MainDb) CreateMessage(group string, name string, age string) error {
	return o.dbSelect(group).InsertRecrods(group, name, age)
}

func (o *MainDb) ReadMessage(group string, name string, age string) error {
	return nil
}
