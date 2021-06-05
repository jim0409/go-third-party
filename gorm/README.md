# 使用GORM連結MySql
1. 使用`docker-compose up -d`快速架設一個mysql伺服器
2. 使用sql指令創造新的usr&db，並且把db權限賦予該usr
```sql
create user 'jim' IDENTIFIED by 'password';
create database `demo_db`;
grant all privileges on demo_db.* to 'jim';
```

# 學習相關紀錄
1. 7/8 實作CRUD
2. 7/14 實作交易

# refer
- golang的orm操作手冊
https://gorm.io/docs/
- gorm tutorial
https://tutorialedge.net/golang/golang-orm-tutorial/
- 使用orm連接到mysql
https://github.com/jinzhu/gorm/issues/403
- mysql創造新的database遇到的坑
https://stackoverflow.com/questions/44916136/error-1064-42000-when-creating-database-in-mysql
- 討論是否每次使用gorm.Open以後要在程序結束以前使用gorm.Close()
https://github.com/jinzhu/gorm/issues/1427
- 如何查看gorm具體執行的sql語句
https://github.com/jinzhu/gorm/issues/1544
- gorm在做(關聯)對多的情況
https://blog.csdn.net/rocky0503/article/details/80915157
- gorm做transaction
https://motion-express.com/blog/gorm:-a-simple-guide-on-crud

# 中文參考
- https://www.bookstack.cn/read/gorm-cn-doc/crud.md
- https://segmentfault.com/a/1190000013216540


# 後記: gorm單元測試的坑...
一開始使用go-mocket(sql-mock)包一層的測試框架，參考
- https://github.com/DATA-DOG/go-sqlmock/issues/118#issuecomment-386692428
發現還是不是很好操作，有些呼叫到底層sql-mock的錯誤，不太好排查...於是打算從sql-mock重新刻一個...但是sql-mock也是會遇到create時發生一些錯誤，於是參考到
- https://github.com/jinzhu/gorm/issues/711#issuecomment-167469666
決定先轉用sqlite3當作測試db(os: 反正就是測試完畢刪除一個檔案哩...)

# 樂觀鎖 vs 悲觀鎖
- https://codertw.com/%E8%B3%87%E6%96%99%E5%BA%AB/121925/
- https://www.itread01.com/content/1533718836.html
- https://segmentfault.com/a/1190000016611415
```
In short, 
悲觀鎖(Pessimistic Locking)：只要開始進行交易，所有對應的表均進行上鎖。禁止任何的改動
樂觀鎖(Optimistic Locking)：會在DB欄位多加上一個`version`代表是否進行該次change或commit，如果version小過目前的col
```