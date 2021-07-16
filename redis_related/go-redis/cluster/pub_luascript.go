package main

import (
	"log"

	"github.com/go-redis/redis"
)

var clusterClient *redis.ClusterClient

func init() {
	clusterClient = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{
			":7000",
			":7001",
			":7002",
			":7003",
			":7004",
			":7005",
		},
	})

	clusterClient.Set("jimkey", "secKey", 0)
}

func main() {
	sha := scriptToCluster(clusterClient)
	ret := clusterClient.EvalSha(sha, []string{"jimkey"}) // not allow []string{} to be `nil` ...
	result, err := ret.Result()
	if err != nil {
		log.Println(err)
	}

	// if result is not 0, ... error in script or invalid command occur
	log.Println(result)
}

func scriptToCluster(c *redis.ClusterClient) string {

	script := redis.NewScript(`
	local key = redis.call("GET", KEYS[1])
	local ok = redis.call("PUBLISH", key, 123)
	if ok then
		return 0
	end
	`)

	var ret string

	c.ForEachMaster(func(m *redis.Client) error {
		if result, err := script.Load(m).Result(); err != nil {
			// panic("缓存脚本到主节点失败")
			log.Println(err)
		} else {
			ret = result
		}
		return nil
	})

	return ret
}
