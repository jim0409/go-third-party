package main

import (
	"fmt"
	"log"

	"github.com/go-redis/redis"
)

func createScript() *redis.Script {
	script := redis.NewScript(`
        local busIdentify   = tostring(KEYS[1])
        local ip            = tostring(KEYS[2])
        local expireSeconds = tonumber(ARGV[1])
        local limitTimes    = tonumber(ARGV[2])
        -- 传入额外参数，请求时间戳
        local timestamp     = tonumber(ARGV[3])
        local lastTimestamp

        local identify  = busIdentify .. "_" .. ip
        local times     = redis.call("LLEN", identify)
        if times < limitTimes then
          redis.call("RPUSH", identify, timestamp)
          return 1
        end

        lastTimestamp = redis.call("LRANGE", identify, 0, 0)
        lastTimestamp = tonumber(lastTimestamp[1])

        if lastTimestamp + expireSeconds >= timestamp then
          return 0
        end

        redis.call("LPOP", identify)
        redis.call("RPUSH", identify, timestamp)

        return 1        
    `)

	return script
}

func scriptCacheToCluster(c *redis.ClusterClient) string {
	script := createScript()
	fmt.Println(script)
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

func main() {
	redisdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{
			":7000",
			":7001",
			":7002",
			":7003",
			":7004",
			":7005",
		},
	})
	// 将脚本缓存到所有节点，执行一次拿到结果即可
	sha := scriptCacheToCluster(redisdb)

	// 执行缓存脚本
	ret := redisdb.EvalSha(sha, []string{
		"limit_vgroup{yes}",
		"172.23.0.1{yes}",
	}, 10, 3, 1548660999)

	if result, err := ret.Result(); err != nil {
		fmt.Println("发生异常，返回值：", err.Error())
	} else {
		fmt.Println("返回值：", result)
	}

	// 示例错误情况，sha 值不存在
	ret1 := redisdb.EvalSha(sha+"error", []string{
		"limit_vgroup{yes}",
		"172.23.0.1{yes}",
	}, 10, 3, 1548660999)

	if result, err := ret1.Result(); err != nil {
		fmt.Println("发生异常，返回值：", err.Error())
	} else {
		fmt.Println("返回值：", result)
	}
}
