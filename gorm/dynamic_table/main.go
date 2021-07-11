package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func hander(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var linkUrl = "jim:password@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=true"

func main() {
	db, err := gorm.Open(mysql.Open(linkUrl), &gorm.Config{})
	hander(err)
	d, err := db.DB()
	hander(err)
	defer d.Close()

	for i := 0; i < 10; i++ {
		tableName := fmt.Sprintf("jim_%d", i)
		db.Exec(DropTable(tableName))
		db.Exec(CreateTable(tableName))

		name := fmt.Sprintf("jim_%d", i)
		age := fmt.Sprintf("%d", i)
		db.Table(tableName).Create(&TableStruct{Name: name, Age: age})
	}

}

type TableStruct struct {
	gorm.Model
	Name string `gorm:"column:name"`
	Age  string `gorm:"column:age"`
}

// Create group id
func CreateTable(name string) string {
	return fmt.Sprintf(`
CREATE TABLE %v (
	id bigint(20) NOT NULL AUTO_INCREMENT,
	name varchar(255) NOT NULL,
	age varchar(255) NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP,
	PRIMARY KEY (ID)
);
`, name)
}

func DropTable(name string) string {
	return fmt.Sprintf("drop table %v;", name)
}
