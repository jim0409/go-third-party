package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.etcd.io/etcd/clientv3"
)

var (
	config clientv3.Config
	client *clientv3.Client
	err    error
)

func main() {
	// 客户端配置
	config = clientv3.Config{
		// Endpoints:   []string{"172.27.43.50:2379"},
		// Endpoints:   []string{"172.30.2.230:2379"},
		Endpoints:   []string{"172.30.2.230:2379", "172.30.1.197:2379", "172.30.3.61:2379"},
		DialTimeout: 5 * time.Second,
	}
	// 建立连接
	// client, err := clientv3.New(config)
	client, err := clientv3.New(config)
	if err != nil {
		log.Fatal(err)
	}
	// key := "/school/class/students"
	key := "/test"

	// write data to etcd
	kv := clientv3.NewKV(client)
	putResp, err := kv.Put(context.TODO(), key, "helios0", clientv3.WithPrevKV())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(putResp.Header.Revision)

	if putResp.PrevKv != nil {
		fmt.Printf("prev Value: %s \n CreateRevision : %d \n ModRevision: %d \n Version: %d \n",
			string(putResp.PrevKv.Value),
			putResp.PrevKv.CreateRevision,
			putResp.PrevKv.ModRevision,
			putResp.PrevKv.Version,
		)
	}

	// read data from etcd
	resget, err := kv.Get(context.TODO(), "/test")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Key %s, Value : %s \n", resget.Kvs[0].Key, resget.Kvs[0].Value)

	resdel, err := kv.Delete(context.TODO(), key)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Key: %v is deleted \n", key)
	fmt.Println(resdel)
}
