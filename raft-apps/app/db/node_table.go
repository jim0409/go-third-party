package db

import "github.com/jinzhu/gorm"

type Node struct {
	gorm.Model
	Port int
	Addr string
}

type ImpNode interface {
	InsertDbRecord(int, string) (uint, error)
	ReturnNodes() (*[]Node, error)
}

var nodeTable = "node_table"

func (n *Node) TableName() string {
	return nodeTable
}

// 在插入新的節點以後，會同時返回該節點的註冊用ID
func (db *Operation) InsertDbRecord(port int, addr string) (uint, error) {
	n := &Node{
		Port: port,
		Addr: addr,
	}
	return n.ID, db.DB.Table(nodeTable).Create(n).Error
}

func (db *Operation) ReturnNodes() (*[]Node, error) {

	ns := &[]Node{}
	err := db.DB.Table(nodeTable).Select(`*`).Where(`deleted_at is NULL`).Scan(ns).Error
	if err != nil {
		return nil, err
	}
	return ns, nil
}
