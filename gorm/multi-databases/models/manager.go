package models

import (
	"fmt"
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

// 紀錄 Group 在哪一個 Node DB 上
type GroupInDB struct {
	gorm.Model
	GroupID string
	NodeID  string
}

// 紀錄 Node DB 的資訊
type NodeInfo struct {
	gorm.Model
	User     string
	Password string
	Type     string
	Database string
	Port     string
	Address  string
}

type MainDb struct {
	DB      MainDB         // 主配置，紀錄 node db 的配置，及動態增加 nodes
	NodeDBs []OPDB         // 紀錄 node db 的實例
	TabeLoc map[string]int // 反查回 table 對應的 node db
}

func (db *Operation) RetriveNodes() ([]NodeInfo, error) {
	nodes := []NodeInfo{}
	if err := db.DB.Table("node_infos").Select("*").Scan(&nodes).Error; err != nil {
		return nil, err
	}
	return nodes, nil
}

func InitMainDB(usr, pwd, tpe, name, port, addr string) DBManager {
	dbc := NewDBConfiguration(usr, pwd, tpe, name, port, addr)
	mdb, err := dbc.NewMainDBConnection()
	if err != nil {
		panic(err)
	}

	err = mdb.migrate(&GroupInDB{}, &NodeInfo{})
	if err != nil {
		panic(err)
	}

	maindb := &MainDb{
		DB:      mdb,
		NodeDBs: make([]OPDB, 0),
		TabeLoc: make(map[string]int),
	}

	defer maindb.UpdateNodes()

	return maindb
}

func (m *MainDb) Closed() error {
	// 要先關閉主要 DB 再關閉節點 DB，避免重複建表
	if err := m.DB.Closed(); err != nil {
		return err
	}

	for _, n := range m.NodeDBs {
		if err := n.Closed(); err != nil {
			return err
		}
	}

	return nil
}

// TODO:考慮是否要定期做排查? 以及增加程序內部緩存?
func (m *MainDb) UpdateNodes() {
	var dbs []OPDB
	cfgs, err := m.DB.RetriveNodes()
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

	m.NodeDBs = dbs
}

type DBManager interface {
	UpdateNodes() // 提供更新 Nodes 資訊
	Closed() error

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
	return strconv.Atoi(id)
}

func (o *MainDb) dbSelect(group string) OPDB {
	if t, ok := o.TabeLoc[group]; ok {
		return o.NodeDBs[t]
	}

	// fmt.Printf("----- no cache with group .. %s -----\n", group)
	loc, err := dbDecision(group)
	if err != nil {
		panic(err)
	}

	o.TabeLoc[group] = loc
	if err := o.NodeDBs[loc].createGroupMsgTabel(group); err != nil {
		panic(err)
	}

	return o.NodeDBs[loc]
}

func (o *MainDb) CreateMessage(group string, name string, age string) error {
	return o.dbSelect(group).InsertRecrods(group, name, age)
}

func (o *MainDb) ReadMessage(group string, name string, age string) error {
	return nil
}
