package main

import (
	"fmt"
	"log"

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
	dtMethod
	Closed() error
	Debug()
}

func (dbc *DBConfig) NewDBConnection() (OPDB, error) {
	db, err := gorm.Open(dbc.DBType, dbc.DBUri)
	if err != nil {
		return nil, err
	}
	db = db.AutoMigrate(&DemoTable{})
	return &OperationDatabase{DB: db}, err
}

func (db *OperationDatabase) Closed() error {
	if err := db.DB.Close(); err != nil {
		return fmt.Errorf("Error happended while closing db : %v\n", err)
	}
	log.Fatalln("Going to close DB")
	return nil
}

// 透過使用Debug()可以轉譯語言為SQL語法
func (db *OperationDatabase) Debug() {
	db.DB = db.DB.Debug()
}

func NewDBConfiguration(user, password string, dbtype string, dbname, dbport string, dbaddress string) *DBConfig {
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
	// if len(os.Args) != 6 {
	// 	fmt.Fprintf(os.Stderr, "Usage: %s <msyql-addr> <mysql-port> <mysql-operation-database> <mysql-user> <mysql-user-password>\n",
	// 		os.Args[0])
	// 	os.Exit(1)
	// }

	// mysqlAddr := os.Args[1]  // 127.0.0.1
	// mysqlPort := os.Args[2]  // 3306
	// mysqlOpDB := os.Args[3]  // mysql
	// mysqlUsr := os.Args[4]   // root
	// mysqUsrPwd := os.Args[5] // secret

	mysqlAddr := "127.0.0.1"
	mysqlPort := "3306"
	mysqlOpDB := "mysql"
	mysqlUsr := "root"
	mysqUsrPwd := "secret"

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
