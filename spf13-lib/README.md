# 前情提要

golang 專案常常見到使用 spf13 相關的套件，其中又以 cobra, viper 最為常見。

- cobra: 主要用於開發指令工具(像是 curl/ ping/ nc)，可以提供一個良好，且不錯的使用者介面
```
(也有一派的人認為 go 的系統缺乏像是 django 那樣 的人機互動，所以引入 cobra ... 但對於 server 做人機互動是否有資安考量，則另外討論...)
```

- viper: 支援多種型態的參數讀取(舉凡 文件/ 傳入參數/ 環境變數)


# 為什麼要探討

雖然 golang 的大多數伺服器開發是不需要使用上述兩個套件包也能完善。

本著增加後續討論的客觀度，以及面對普遍公司的使用套件包習慣做練習


# refer:
- https://github.com/spf13
- https://github.com/spf13/cobra
- https://github.com/spf13/viper
