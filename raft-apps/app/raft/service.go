package raft

import (
	"go.etcd.io/etcd/raft/v3/raftpb"
)

type RaftNode struct {
	id       int
	join     bool
	kvport   int
	clusters []string

	proc  chan string
	confc chan raftpb.ConfChange
}

func (r *RaftNode) RunRaftNode() {
	var kvs *kvstore
	getSnapshot := func() ([]byte, error) { return kvs.getSnapshot() }

	// id 使用 uint 應該沒問題，要修改 newRaftNode
	commitC, errorC, snapshotterReady := newRaftNode(r.id, r.clusters, r.join, getSnapshot, r.proc, r.confc)

	kvs = newKVStore(<-snapshotterReady, r.proc, commitC, errorC)

	// the key-value http handler will propose updates to raft
	serveHttpKVAPI(kvs, r.kvport, r.confc, errorC)
}

func (r *RaftNode) Close() {
	close(r.proc)  // prposeC
	close(r.confc) // confChangeC
}

func InitRaftNode(id int, kvport int, clusters []string, join bool) *RaftNode {
	return &RaftNode{
		id:       id,
		kvport:   kvport,
		clusters: clusters,
		join:     join,
		proc:     make(chan string),
		confc:    make(chan raftpb.ConfChange),
	}
}
