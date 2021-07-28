package main

import (
	"flag"
	"strings"

	"go-third-party/raft-apps/app/db"

	"go.etcd.io/etcd/raft/v3/raftpb"
)

/*
1. start service with InitNodeConfig
2. throw input variables ndoe_addr, node_port into mysql
3. query mysql records, if records length is greater than 1, add join(true) flag
4. else return `id` as the started node `id`
*/

func InitDB() (db.OPDB, error) {
	mysqlAddr := "mysql"
	mysqlPort := "3306"
	mysqlOpDB := "raft"
	mysqlUsr := "raft"
	mysqUsrPwd := "raft"

	return db.NewDBConfiguration(mysqlUsr, mysqUsrPwd, "mysql", mysqlOpDB, mysqlPort, mysqlAddr).NewDBConnection()
}

func InitNodeConfig() (int, int, string, bool) {
	id := flag.Int("id", 0, "node ID") // 主要用於 recover status

	// 對於初次註冊的節點，需要帶入一些參數
	// TODO: 給予系統話參數，或設定檔考慮 e.g. randomPort
	port := flag.Int("port", 0, "key-value server port")
	addr := flag.String("addr", "http://127.0.0.1:12379", "used for peer-connection  port")
	join := flag.Bool("join", false, "join an existing cluster")
	flag.Parse()

	db, err := InitDB()
	if err != nil {
		panic(err)
	}

	if *id == 0 {
		// 自動加入新節點
		aid, err := db.InsertDbRecord(*port, *addr)
		if err != nil {
			panic(err)
		}
		*id = aid

	}

	if *port == 0 {
		nodeInfo, err := db.ReturnNodeInfo(*id)
		if err != nil {
			panic(err)
		}
		*port = nodeInfo.Port
	}

	clusters, err := db.GetClusterIps()
	if err != nil {
		panic(err)
	}

	if len(clusters) > 1 {
		*join = true
	}
	ips := strings.Join(clusters, ",")

	return *id, *port, ips, *join
}

func main() {

	id, kvport, cluster, join := InitNodeConfig()

	proposeC := make(chan string)
	defer close(proposeC)
	confChangeC := make(chan raftpb.ConfChange)
	defer close(confChangeC)

	var kvs *kvstore
	getSnapshot := func() ([]byte, error) { return kvs.getSnapshot() }
	clusters := strings.Split(cluster, ",")

	// id 使用 uint 應該沒問題，要修改 newRaftNode
	commitC, errorC, snapshotterReady := newRaftNode(id, clusters, join, getSnapshot, proposeC, confChangeC)

	kvs = newKVStore(<-snapshotterReady, proposeC, commitC, errorC)

	// the key-value http handler will propose updates to raft
	serveHttpKVAPI(kvs, kvport, confChangeC, errorC)
	// handler.addNode()
	for {

	}

}
