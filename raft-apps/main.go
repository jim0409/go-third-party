package main

import (
	"flag"
	"fmt"
	"strings"

	"go.etcd.io/etcd/raft/v3/raftpb"
)

func InitRedis(addr string, password string) redisDAO {
	return NewRedisClient(addr, password)
}

func InitNodeConfig() (int, int, string, bool) {
	id := flag.Int("id", 1, "node ID")
	kvport := flag.Int("port", 9121, "key-value server port")
	addr := flag.String("cluster", "http://127.0.0.1:9021", "comma separated cluster peers")
	join := flag.Bool("join", false, "join an existing cluster")
	flag.Parse()

	return *id, *kvport, *addr, *join
}

func registNode(id, port int, addr string, join bool) error {
	rds := InitRedis("127.0.0.1:6379", "yourpassword")
	strs, err := rds.lrange("cluster")
	if err != nil {
		return err
	}
	fmt.Println(strs)
	// host, err := os.Hostname()
	// if err != nil {
	// 	return err
	// }

	// prefix as random generate id? -- no, due to lrange would auto gen `id`
	// prefix := fmt.Sprintf("%v_%v", time.Now().Nanosecond(), host)
	record := fmt.Sprintf("%v:%v", addr, port)
	if err := rds.lpush("cluster", record); err != nil {
		return err
	}
	// _ = rds

	return nil
}

func main() {

	proposeC := make(chan string)
	defer close(proposeC)
	confChangeC := make(chan raftpb.ConfChange)
	defer close(confChangeC)

	id, kvport, cluster, join := InitNodeConfig()
	var kvs *kvstore
	getSnapshot := func() ([]byte, error) { return kvs.getSnapshot() }
	clusters := strings.Split(cluster, ",")

	commitC, errorC, snapshotterReady := newRaftNode(id, clusters, join, getSnapshot, proposeC, confChangeC)

	kvs = newKVStore(<-snapshotterReady, proposeC, commitC, errorC)

	// the key-value http handler will propose updates to raft
	handler := serveHttpKVAPI(kvs, kvport, confChangeC, errorC)
	// handler.addNode()
	_ = handler

}
