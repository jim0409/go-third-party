# Introduction
交易，一種非0即1的行為。符合ACIID

# 摘自GORM的文檔
GORM默認會將單個的`create`/`update`/`delete`操作封裝在交易內進行處理，以確保數據的完整性。

```
transaction
# 開啟交易模式
tx := db.Begin

# 在交易模式中執行劇裡的資料庫操作(noted: 交易下，使用的物件是'tx'而非'db')
tx.Create(...)

# 如果發生錯誤則執行滾回
tx.Rollback()

# 確認交易成功(未發生錯誤)，執行提交
tx.Commit()

```

# refer
- 交易
  - http://karenten10-blog.logdown.com/posts/192629-database-transaction-1-acid
- 套件
  - http://gorm.io/zh_CN/docs/transactions.html
- 學習範例
  - http://hopehook.com/2017/08/21/golang_transaction/
  - https://stackoverflow.com/questions/16184238/database-sql-tx-detecting-commit-or-rollback