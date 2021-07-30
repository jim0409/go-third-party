package raft

import (
	"go.etcd.io/etcd/raft/v3/raftpb"
)

type RaftNode struct {
	id       int
	join     bool
	kvport   int
	clusters []string
}

func (r *RaftNode) RunRaftNode() {
	proposeC := make(chan string)
	confChangeC := make(chan raftpb.ConfChange)

	var kvs *kvstore
	getSnapshot := func() ([]byte, error) { return kvs.getSnapshot() }

	// id 使用 uint 應該沒問題，要修改 newRaftNode
	commitC, errorC, snapshotterReady := newRaftNode(r.id, r.clusters, r.join, getSnapshot, proposeC, confChangeC)

	kvs = newKVStore(<-snapshotterReady, proposeC, commitC, errorC)

	// the key-value http handler will propose updates to raft
	serveHttpKVAPI(kvs, r.kvport, confChangeC, errorC)
}

// int, int, string, bool
func InitRaftNode(id int, kvport int, clusters []string, join bool) *RaftNode {
	return &RaftNode{
		id:       id,
		kvport:   kvport,
		clusters: clusters,
		join:     join,
	}
}
