# 介紹
cobra, 多數來自全端或前端的工程師愛用的指令包之一


# 歷史因素
1. 大多數方便框架, 舉例 django/ flask/ vue/ react/ beego/ echo ...(支持 quick start 的框架建構模式)
2. 同時也支持部分的除錯指令(包含 go 本身也有 go run/test/get/mod args)
3. 基於此種因素，工程師會在開發前期評估是否要加入例如 nginx -s reload 之類的指令?而考慮服務支援到指令集


# 應用場景
透過 cobra 執行不同的模組


# 優點
簡化指令生成，方便"新手"快速開發


# 缺點
因顧及到生成指令代碼，前期需要謹慎規劃。避免後續耦合過高以至於專案難以維護 ..



# 參照
- https://github.com/spf13/cobra
- http://liuqh.icu/2021/11/07/go/package/28-cobra/
- https://darjun.github.io/2020/01/17/godailylib/cobra/


# 題外話
不要看到 cobra 被 docker/ etcd/ kubectl 引用就覺得專案一定也要來用這個
1. 就早期的專案架構，比較傾向 server 端以外再另外出一個 sdk 給 client 端作使用 e.g. vmware
2. 即便是 docker/ etcd/ kubectl 都有使用到本地持久化的設定檔配置，對於一般(微)業務型的服務可能就不是一個太好的選擇

