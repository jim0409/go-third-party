# intro

延遲實作，動態生成 mysql table

# 應用場景

單張大表持續有資料寫入 .. 但是每次訪問查詢，並不需要關聯


# refer:
<!--
2021/07/11 .. 目前 gorm 2.0 還沒有釋出 置換 TableName 方法來達到動態增加 table的功能
-->
- https://gorm.io/docs/migration.html
- https://gorm.io/zh_CN/docs/conventions.html
- https://dev.mysql.com/doc/refman/8.0/en/create-table.html


### 命名規則
- https://gorm.io/docs/gorm_config.html
- https://gorm.io/zh_CN/docs/conventions.html
