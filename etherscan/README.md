## intro

etherscan api practice

### api export
1. 檢查指定帳戶下是否有其對應交易紀錄


#### test script

```bash
curl -XPOST http://127.0.0.1:8000/etherscan -d '{"account":"0x755229B87cEc1d55c8b1Aa0c66C1be4D8136E8CF", "targets": ["0x5f432212023ecd0019aca196d561e9ac96a95dee5f38ddfeae89647df45d6126"]}'
```

### refer:
- https://ethereum.org/zh/developers/docs
- https://etherscan.io/
- https://learnblockchain.cn/docs/etherscan
- http://cw.hubwiz.com/card/c/etherscan-api
