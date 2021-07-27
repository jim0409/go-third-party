#!/bin/bash
cmd=$1

function start_node1() {
	echo "---- start node1 ----"
	go run *.go --id 1 --cluster http://127.0.0.1:12379 --port 12380
}

function start_node2() {
	echo "---- start node2 ----"
	go run *.go --id 2 --cluster http://127.0.0.1:12379,http://127.0.0.1:22379 --port 22380 --join
}

function start_node3() {
	echo "---- start node3 ----"
	#<!> can put partial cluster ips, or cause err to down
	# go run *.go --id 3 --cluster http://127.0.0.1:12379,http://127.0.0.1:32379 --port 32380 --join
	go run *.go --id 3 --cluster http://127.0.0.1:12379,http://127.0.0.1:22379,http://127.0.0.1:32379 --port 32380 --join
}

function start_node4() {
	echo "---- start node4 ----"
	go run *.go --id 4 --cluster http://127.0.0.1:12379,http://127.0.0.1:22379,http://127.0.0.1:32379,http://127.0.0.1:42379 --port 42380 --join
}

function add_node_2() {
	echo "---- add node 2 ----"
	curl -L http://127.0.0.1:12380/2 -XPOST -d http://127.0.0.1:22379
}

function add_node_3() {
	echo "---- add node 3 ----"
	curl -L http://127.0.0.1:12380/3 -XPOST -d http://127.0.0.1:32379
}

function add_node_4() {
	echo "---- add node 4 ----"
	curl -L http://127.0.0.1:12380/4 -XPOST -d http://127.0.0.1:42379
}

function del_node_2() {
	echo "---- del node 2 ----"
	curl -L http://127.0.0.1:12380/2 -XDELETE
}

function del_node_3() {
	echo "---- del node 3 ----"
	curl -L http://127.0.0.1:12380/3 -XDELETE
}

function del_node_4() {
	echo "---- del node 4 ----"
	curl -L http://127.0.0.1:12380/4 -XDELETE
}

function add_key_value() {
	local key=$1
	local value=$2
	curl -XPUT http://127.0.0.1:12380/$key -d "$value"
	echo "---- add key: $key, value: $value ----"
}

function get_key() {
	local key=$1
	value=`curl -XGET http://127.0.0.1:12380/$key` 
	echo "---- s1 => get key: $key, value: $value ----"
	value=`curl -XGET http://127.0.0.1:22380/$key` 
	echo "---- s2 => get key: $key, value: $value ----"
	value=`curl -XGET http://127.0.0.1:32380/$key` 
	echo "---- s3 => get key: $key, value: $value ----"
	value=`curl -XGET http://127.0.0.1:42380/$key` 
	echo "---- s4 => get key: $key, value: $value ----"
}

function del_key() {
	local key=$1
	value=`curl -XDELETE http://127.0.0.1:12380/$key`
	echo "---- del key: $key ----"
}

case $cmd in
	"s1")
	start_node1
;;
	"s2")
	start_node2
;;
	"s3")
	start_node3
;;
	"s4")
	start_node4
;;
	"a2")
	add_node_2
;;
	"a3")
	add_node_3
;;
	"a4")
	add_node_4
;;
	"d2")
	del_node_2
;;
	"d3")
	del_node_3
;;
	"d4")
	del_node_4
;;
	"add")
	add_key_value $2 $3
;;
	"get")
	get_key $2
;;
	"cls")
	rm -rf raftexample-*
;;
	*)
	echo "no such command support"
esac
