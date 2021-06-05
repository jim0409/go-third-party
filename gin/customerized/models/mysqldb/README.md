# Many To Many
Many to Many adds a join table between two models.

For example, if your application includes users and languages,

and a user can speak many languages, and many users can speak a specfied language.

```go
// User has and belongs to many languages, use `user_languages` as join table
type User struct {
  gorm.Model
  Languages         []*Language `gorm:"many2many:user_languages;"`
}

type Language struct {
  gorm.Model
  Name string
  Users         	  []*User     `gorm:"many2many:user_languages;"`
}
```

create some record for the many-to-many tables
```go
var users []User
language := Language{}

db.First(&language, "id = ?", 111)

db.Model(&language).Related(&users,  "Users")

//// SELECT * FROM "users" INNER JOIN "user_languages" ON "user_languages"."user_id" = "users"."id" WHERE  ("user_languages"."language_id" IN ('111'))
```


# refer:
- http://gorm.io/docs/many_to_many.html
- https://github.com/jinzhu/gorm/issues/754