package mysqldb

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/mattn/go-sqlite3"
)

var (
	mysqlDbConfig *MySQLConfig
	mysqlDb       MySQLDBAccessObject
	once          sync.Once
)

type MySQLConfig struct {
	DBName           string
	DBHost           string
	DBPort           string
	DBUsr            string
	DBPassword       string
	DBLogEnable      bool
	DBMaxConnection  int
	DBIdleConnection int
	DBUri            string
}

func LoadMySQLDBConfig(dbName, dbHost, dbPort, dbUsr, dbPassword string, dbLogEnable bool, dbMaxConnection, dbIdleConnection int) {
	mysqlDbConfig = &MySQLConfig{
		DBName:           dbName,
		DBHost:           dbHost,
		DBPort:           dbPort,
		DBUsr:            dbUsr,
		DBPassword:       dbPassword,
		DBLogEnable:      dbLogEnable,
		DBMaxConnection:  dbMaxConnection,
		DBIdleConnection: dbIdleConnection,
		DBUri: fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			dbUsr, dbPassword, dbHost, dbPort, dbName,
		),
	}
}

func RetriveMySQLDBAccessObj() MySQLDBAccessObject {
	once.Do(func() {
		// mysqlDb = mysqlDb
		mysqlDb = &mysqlDBObj{}
	})
	return mysqlDb
}

func StartMySQLDB() error {
	var err error
	mysqlDb = RetriveMySQLDBAccessObj()
	mysqlDb, err = initMySqlDB(mysqlDbConfig)
	return err
}

type mysqlDBObj struct {
	DB *gorm.DB
}

func (db *mysqlDBObj) Close() error {
	return db.DB.Close()
}

func (db *mysqlDBObj) ExecSql(sqlStr string) error {
	var err error
	for i := 0; i < 5; i++ {
		err = db.DB.Exec(sqlStr).Error
		if err != nil {
			if i < 4 {
				time.Sleep(time.Duration((i + 1)) * 2 * time.Second)
			}
			continue
		} else {
			break
		}
	}
	if err != nil {
		return fmt.Errorf("insert sql=%v to dbs err: %v", sqlStr, err)
	}

	return nil
}

type MySQLDBAccessObject interface {
	migration(interface{})
	Close() error
	ExecSql(string) error
	CreateRecord(string, ...string) error
	QueryRecord(string) (*[]Language, error)
}

func initMySqlDB(c *MySQLConfig) (MySQLDBAccessObject, error) {
	var db *gorm.DB
	var err error

	dbtype := "mysql"
	if os.Getenv("TestDB") == "test" {
		dbtype = "sqlite3"
		c.DBUri = "/tmp/gorm.db"
	}

	if db, err = gorm.Open(dbtype, c.DBUri); err != nil {
		return nil, fmt.Errorf("Connection to MySQL DB error : %v", err)
	}

	db.DB().SetMaxOpenConns(c.DBMaxConnection)
	db.DB().SetMaxIdleConns(c.DBIdleConnection)

	if err = db.DB().Ping(); err != nil {
		return nil, fmt.Errorf("Ping MySQL db error : %v", err)
	}

	db.LogMode(c.DBLogEnable)
	db.SingularTable(true)

	return &mysqlDBObj{DB: db}, nil
}

/*
	add another registry to migrate all the db at once instead of migration one by one
	may claim a map and return table access object
*/
func (db *mysqlDBObj) migration(v interface{}) {
	db.DB.AutoMigrate(v)
}
