# intro
使用mongodb
1. write/get/update/delete data
2. transaction


# 假設有一個目標叫做`tags`的collection
### example for mongodump data
> mongodump -u root -p password --authenticationDatabase admin -d tags -o /tmp/mongo_tags

### example for restore data
> mongorestore -u root -p password --authenticationDatabase admin -d tags ./tags/

# mongo connections pool
- https://www.jianshu.com/p/d50559e8e10c
- https://www.jianshu.com/p/3fbdae6c364a
- https://github.com/kmnemon/goproject/blob/master/golang-mongodb-pool-master/db_connection_pool.go



# refer:
starting and setup
- https://www.mongodb.com/blog/post/quick-start-golang--mongodb--starting-and-setup


package reference
- https://github.com/mongodb/mongo-go-driver


mongodb example
- https://www.mongodb.com/blog/post/quick-start-golang--mongodb--how-to-create-documents


how to use mongo with golang
- https://kb.objectrocket.com/mongo-db/how-to-get-mongodb-documents-using-golang-446