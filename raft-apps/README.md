# intro
利用 /etcd/raft 包，製作一個 app ~

# goal
via raft lib to build some apps

# quick start
1. 啟動 node1
> ./run s1

2. 啟動以及增加 node2
> ./run a2 &&  ./run s2

3. 啟動以及增加 node3
> ./run a3 && ./run s3

4. 啟動以及增加 node4
> ./run a4 && ./run s4

5. 刪除 node4
> ./run d4 (shutdown$ ./run s4)
<!--
1. 可以反覆驗證，刪除其他節點
2. 已知，刪除節點後。無法再加入重複id節點。只能嚴格遞增`id`
-->


# 增加key-value
1. 增加一個 key: jim, value: weng 的鍵值
> ./run add jim weng


# 驗證key-valu
1. 取得 key: jim 的值
> ./run get jim


# 備註
1. 當節點故障後重新啟動
> 因為有做`snapshot`所以可以直接啟動，不需要更動任何資料

2. 清除所有資料
> ./run cls



# refer:
- https://github.com/etcd-io/etcd/tree/main/raft
- https://github.com/etcd-io/etcd/tree/main/contrib/raftexample