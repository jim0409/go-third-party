package raft

import (
	"fmt"
	"net/http"
	"strings"

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

	if r.join {
		r.regist()
	}

	// the key-value http handler will propose updates to raft
	serveHttpKVAPI(kvs, r.kvport, r.confc, errorC)
}

func (r *RaftNode) Close() {
	close(r.proc)  // prposeC
	close(r.confc) // confChangeC
	// r.unregist() // may consider to unregistr ?
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

func (r *RaftNode) regist() {
	// params := url.Values{}
	// params.Add("http://127.0.0.1:22379", "")
	// body := strings.NewReader(params.Encode())
	peeradrr := r.clusters[len(r.clusters)-1]
	// fmt.Println("----- peer addr --------- ", peeradrr)
	// body := strings.NewReader("http://127.0.0.1:22379")
	body := strings.NewReader(peeradrr)
	// fmt.Println("----- boody --------- ", body)

	// url := fmt.Sprintf("%v/%d", r.clusters[0], r.id)
	// req, err := http.NewRequest("POST", url, body)
	// req, err := http.NewRequest("POST", "http://127.0.0.1:12380/2", body)
	req, err := http.NewRequest("POST", fmt.Sprintf("http://127.0.0.1:12380/%d", r.id), body)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}
