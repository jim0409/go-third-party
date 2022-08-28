# advance

測試 FSM 的一些功能
1. 狀態函數
> go test -v -run TestDoorFsm

2. 自動生成流程代碼
> go test -v -run TestFsmVisual

3. 效能測試
> go test -v -bench=. -run=BenchmarkDoorFsm
```log
goos: darwin
goarch: amd64
pkg: go-third-party/fsm/door
cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
BenchmarkDoorFsm
BenchmarkDoorFsm-16      1148785              1019 ns/op
ok      go-third-party/fsm/door 2.317s
```