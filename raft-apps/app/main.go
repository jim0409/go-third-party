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
	mysqlAddr := "127.0.0.1"
	mysqlPort := "3306"
	mysqlOpDB := "raft"
	mysqlUsr := "raft"
	mysqUsrPwd := "raft"

	return db.NewDBConfiguration(mysqlUsr, mysqUsrPwd, "mysql", mysqlOpDB, mysqlPort, mysqlAddr).NewDBConnection()
}

func InitNodeConfig() (int, int, string, bool) {
	id := flag.Int("id", 1, "node ID")
	kvport := flag.Int("port", 9121, "key-value server port")
	cluster := flag.String("cluster", "http://127.0.0.1:9021", "comma separated cluster peers")
	join := flag.Bool("join", false, "join an existing cluster")
	flag.Parse()
	return *id, *kvport, *cluster, *join
}

func main() {

	id, kvport, cluster, join := InitNodeConfig()
	// db, err := InitDB()
	// if err != nil {
	// 	panic(err)
	// }

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
