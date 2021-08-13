package models

import (
	"fmt"

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
type GroupInNodes struct {
	gorm.Model
	GroupID  string `gorm:"unique"`
	NodeID   int
	NodeInfo NodeInfo `gorm:"foreignKey:NodeID"`
	// NodeID  NodeInfo `gorm:"foreign` // foreign key of node_info_id
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
	SettingDB MainDB         // 主配置，紀錄 node db 的配置，及動態增加 nodes
	NodeDBs   []OPDB         // 紀錄 node db 的實例
	TabeLoc   map[string]int // 緩存; 回傳 group 對應的 node db --TODO: 改成併發寫安全，讀不設限的緩存套件
}

var (
	groupIndNodesTableName = "group_in_nodes"
	nodeInfoTableName      = "node_infos"
)

func (n *GroupInNodes) TableName() string {
	return groupIndNodesTableName
}

func (n *NodeInfo) TableName() string {
	return nodeInfoTableName
}

// 實現 MainDB interface { ... }
func (o *Operation) RetriveNodes() ([]NodeInfo, error) {
	nodes := []NodeInfo{}
	if err := o.DB.Table(nodeInfoTableName).Select("*").Scan(&nodes).Error; err != nil {
		return nil, err
	}
	return nodes, nil
}

func (o *Operation) AddNodeInfos(ns []NodeInfo) error {
	return o.DB.Create(ns).Error
}

func (o *Operation) AddGroupInNodes(groupId string, nodeId int) error {
	return o.DB.Create(&GroupInNodes{
		GroupID: groupId,
		NodeID:  nodeId,
	}).Error
}

func (o *Operation) QueryGroupLoc(groupId string) (int, error) {
	loc := -1 // 預設 loc 為 -1，表示查找不到
	return loc, o.DB.Table(groupIndNodesTableName).Select(`node_id`).Where("group_id = ?", groupId).Scan(&loc).Error
}

// NodeStatic 回傳當下節點的使用狀況
func (o *Operation) NodeStatic() (int, error) {
	var loc, biggest int
	var err error
	var node_ids []int

	// db 只把 node_id 查出來就好，後面的統計用 golang 執行
	if err := o.DB.Table(groupIndNodesTableName).Select(`node_id`).Scan(&node_ids).Error; err != nil {
		return loc, err
	}

	m := make(map[int]int, len(node_ids))
	for _, id := range node_ids {
		m[id]++
		if m[id] > biggest {
			loc, biggest = id, m[id]
		}
	}

	return loc, err
}

func InitMainDB(usr, pwd, tpe, name, port, addr string) DBManager {
	dbc := NewDBConfiguration(usr, pwd, tpe, name, port, addr)
	mdb, err := dbc.NewMainDBConnection()
	if err != nil {
		panic(err)
	}

	err = mdb.migrate(&GroupInNodes{}, &NodeInfo{})
	if err != nil {
		panic(err)
	}

	maindb := &MainDb{
		SettingDB: mdb,
		NodeDBs:   make([]OPDB, 0),
		TabeLoc:   make(map[string]int),
	}

	defer maindb.UpdateNodes()

	return maindb
}

func (m *MainDb) Closed() error {
	// 要先關閉主要 DB 再關閉節點 DB，避免重複建表
	if err := m.SettingDB.Closed(); err != nil {
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
	cfgs, err := m.SettingDB.RetriveNodes()
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
	fmt.Println(m.NodeDBs)
}

type DBManager interface {
	UpdateNodes() // 提供更新 Nodes 資訊
	Closed() error

	DBWriter
	DBReader
	DBBroker
}

type DBWriter interface {
	CreateMessage(string, string, string) error
}

type DBReader interface {
	ReadMessage(string, []string) ([]map[string]interface{}, error)
}

type DBBroker interface {
	PurgeCache() error
}

func (m *MainDb) PurgeCache() error {
	newTable := make(map[string]int)
	m.TabeLoc = nil
	m.TabeLoc = newTable
	return nil
}

// TODO: 製作決策器 .. 要下決策決定他是去哪一個 shard
/*
	1. 獲取既有 DB 的資料
	2. 獲取 GroupInDB 分配數量
	3. 分配 group 至 GroupInDB 數量最少的 DB
	4. [交易]寫入 GroupInDB 紀錄，成功後再返回 DB 編號
*/
func (m *MainDb) locDecide(group string) (int, error) {
	nid, err := m.SettingDB.QueryGroupLoc(group)
	if err == nil {
		return nid, nil
	}

	// 假設錯誤是 node_id not found
	return m.SettingDB.NodeStatic()
}

func (m *MainDb) dbSelect(group string) (OPDB, error) {
	// 如果本地緩存存在該 group 寫入位置，直接回傳對應位置
	if t, ok := m.TabeLoc[group]; ok {
		return m.NodeDBs[t], nil
	}

	// fmt.Printf("----- no cache with group .. %s -----\n", group)
	loc, err := m.locDecide(group)
	loc = loc - 1
	if err != nil {
		return nil, err
	}

	if err := m.NodeDBs[loc].createGroupMsgTabel(group); err != nil {
		return nil, err
	}
	// 確定 create table 成功後，才將 loc 位置放入 group
	m.TabeLoc[group] = loc

	return m.NodeDBs[loc], nil
}

func (m *MainDb) CreateMessage(group string, name string, age string) error {
	db, err := m.dbSelect(group)
	if err != nil {
		return err
	}
	return db.InsertRecrods(group, name, age)
}

func (m *MainDb) ReadMessage(group string, option []string) ([]map[string]interface{}, error) {
	db, err := m.dbSelect(group)
	if err != nil {
		return nil, err
	}
	return db.QueryTable(group, option)
}
