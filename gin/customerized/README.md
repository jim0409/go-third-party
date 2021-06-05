# customerized
why customerized?
因應目前http框架，的一些客制需求，自己打造一個自己比較順手的組合框架

# 滿足
1. 使用workers pool: 便利運維人員控管服務資源
2. 支援httpClient的客制化: 有效限制http連線池的總量，及客製化的代理IP轉跳
3. middleware開發: 編寫好自己的middleware放在controller下面
4. 使用簽名機制: 可以友善開發者在測試或者對接API時的方便性


### 串接workers pool與golang MVC框架gin
request -> gin -> gin-handler -> requestTask -> workers pool -> response


### 為什麼要這麼做?
gin已經是一個很成熟的框架了，本質上並不需要去實作worker pools去增加請求處理效能。相反的，透過workers pool反而可能導致效能降低？

<!-- 
### 哪種情境下需要使用?
1. 爬蟲，控制爬蟲池的工作者爬蟲，避免請求量過大時會耗費過多資源(透過改變工作池的大小可以限制資源，及集中管理錯誤訊息)
2. 金流，控制交易速度。透過Ticker可以去限制工作池內的接案狀況，避免過於頻繁的請求直接湧入後台。
3. 需要監控效能的服務，透過監控工作池的活動情形。可以限制服務的使用資源。便於放入像是k8s或是資源有限的環境
-->

### 模組化簽名功能
便利於一些需要文件上加簽的情境




# refer:
- https://stackoverflow.com/questions/36434332/does-gin-gonic-process-requests-in-parallel
- https://medium.com/@pulumati.priyank/go-web-programming-mvc-architecture-based-web-app-73efdb826aa1