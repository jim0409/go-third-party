# intro

使用 golang plugin 套件


# folder arch
```
├── main.go
└── plugin
    └── plugin1.go      // 用來實現的 plugin
```

# quick start
1. build plugin.so
<!-- > go build --tags plugin -o plugin.so . -->
> go build -buildmode=plugin -o plugin.so plugin.go

2. go run with release tags
> go run --tags release .

# refer:
- https://www.gushiciku.cn/pl/gFj1/zh-tw
