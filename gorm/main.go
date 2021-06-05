package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type DBConfig struct {
	User      string
	Password  string
	DBType    string
	DBName    string
	DBAddress string
	DBPort    string
	DBUri     string
}

type OperationDatabase struct {
	DB *gorm.DB
}

type OPDB interface {
	create(name string, email string) error
	queryWithName(name string) (string, error)
	updateEmail(name string, email string) error
	deleteData(name string, email string) error
	Closed() error
	debug()
	transaction() error
}

func (dbc *DBConfig) NewDBConnection() (OPDB, error) {
	// connection :=
	db, err := gorm.Open(dbc.DBType, dbc.DBUri)
	if err != nil {
		return nil, err
	}
	db = db.AutoMigrate(&DemoTable{})
	return &OperationDatabase{DB: db}, err
}

func NewDBConfiguration(user string, password string, dbtype string, dbname string, dbport string, dbaddress string) *DBConfig {
	return &DBConfig{
		User:      user,
		Password:  password,
		DBType:    dbtype,
		DBName:    dbname,
		DBPort:    dbport,
		DBAddress: dbaddress,
		DBUri:     user + ":" + password + "@tcp(" + dbaddress + ":" + dbport + ")/" + dbname + "?charset=utf8&parseTime=True&loc=Local",
	}
}

func main() {
	fmt.Println("Go ORM Tutorial")
	if len(os.Args) != 6 {
		fmt.Fprintf(os.Stderr, "Usage: %s <msyql-addr> <mysql-port> <mysql-operation-database> <mysql-user> <mysql-user-password>\n",
			os.Args[0])
		os.Exit(1)
	}

	mysqlAddr := os.Args[1]  // 127.0.0.1
	mysqlPort := os.Args[2]  // 3306
	mysqlOpDB := os.Args[3]  // mysql
	mysqlUsr := os.Args[4]   // root
	mysqUsrPwd := os.Args[5] // secret

	newDB := NewDBConfiguration(mysqlUsr, mysqUsrPwd, "mysql", mysqlOpDB, mysqlPort, mysqlAddr)
	db, err := newDB.NewDBConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Closed()

	if err := db.create("jim", "test-mail"); err != nil {
		fmt.Printf("Error happend while creating records %s\n", err)
	} else {
		fmt.Println("success insert record into mysql")
	}

	if queryrest, err := db.queryWithName("jim"); err != nil {
		fmt.Printf("Error happend while querying %s\n", err)
	} else {
		fmt.Printf("The query result is %v\n", queryrest)
	}

	if err := db.cleanAll(); err != nil {
		fmt.Printf("Error happend while cleaning %s\n", err)
	} else {
		fmt.Println("success drop mysql talbe demo_tables")
	}

}
