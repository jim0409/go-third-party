package main

import (
	"fmt"
	"math/rand"

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
	DB      OPDB           // 主配置，紀錄 shard dbs 的配置，及動態增加 shards
	ShardDb []OPDB         // 紀錄 shard dbs 的實例
	TabeLoc map[string]int // 反查回 shard db 的位置
}

func InitMainDB(usr, pwd, tpe, name, port, addr string) DBManager {
	mamindb := NewDBConfiguration(usr, pwd, tpe, name, port, addr)
	mdb, err := mamindb.NewDBConnection()
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

/*
TODO:
	改成 query main db 下的 ShardInfo 表
	考慮是否要定期做排查?以及增加程序內部緩存?
*/
func (m *MainDb) retriveDBs() [][]string {
	dbs := [][]string{
		{"jim", "password", "mysql", "message", "3301", "127.0.0.1"},
		{"jim", "password", "mysql", "message", "3302", "127.0.0.1"},
		{"jim", "password", "mysql", "message", "3303", "127.0.0.1"},
	}

	return dbs
}

func (m *MainDb) StartShardDbs() {
	var dbs []OPDB
	cfgs := m.retriveDBs()
	for _, cfg := range cfgs {
		obj := NewDBConfiguration(cfg[0], cfg[1], cfg[2], cfg[3], cfg[4], cfg[5])
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

func randnum() int {
	return rand.Intn(1)
}

func (o *MainDb) dbSelect(group string) OPDB {
	if t, ok := o.TabeLoc[group]; ok {
		return o.ShardDb[t]
	}

	var loc = randnum() // TODO: .. 要下決策決定他是去哪一個 shard
	o.TabeLoc[group] = loc
	if err := o.ShardDb[loc].createGroupMsgTabel(group); err != nil {
		panic(err)
	}
	fmt.Println(loc)

	return o.ShardDb[loc]
}

func (o *MainDb) CreateMessage(group string, name string, age string) error {
	return o.dbSelect(group).InsertRecrods(group, name, age)
}

func (o *MainDb) ReadMessage(group string, name string, age string) error {
	return nil
}
