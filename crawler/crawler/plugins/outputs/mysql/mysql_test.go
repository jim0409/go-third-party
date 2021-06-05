package mysql

import (
	"log"
	"os"
	"testing"

	"github.com/jimweng/crawler/crawler/utils"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

var demo_pts = []*utils.PKGContent{
	&utils.PKGContent{Name: "name1", Parent: "parent1", Synopsis: "synopsis1", Href: "href1"},
	&utils.PKGContent{Name: "name2", Parent: "parent2", Synopsis: "synopsis2", Href: "href2"},
	&utils.PKGContent{Name: "name3", Parent: "parent3", Synopsis: "synopsis3", Href: "href3"},
	&utils.PKGContent{Name: "name1", Parent: "parentu1", Synopsis: "synopsisu1", Href: "hrefu1"},
}

func init() {
	clearTestEnv()
}

// Set mock DB env
func mockTestEnv() (dbOperationInterface, error) {
	var sqlc = SQLConfig{DBType: "sqlite3"}
	connection := "/tmp/gorm.db"
	sqlc.ConnectionUrl = connection
	opdb, err := sqlc.newDBConnection()

	return opdb, err
}

// Clear origin test env if existed
func clearTestEnv() {
	log.Println("Clear test env")
	if _, err := os.Stat("/tmp/gorm.db"); !os.IsNotExist(err) {
		log.Println("Remove Origin gorm.db Files")
		os.Remove("/tmp/gorm.db")
	}
	if _, err := os.Stat("/tmp/mocksqlite3.db"); !os.IsNotExist(err) {
		log.Println("Remove Origin mocksqlite3 Files")
		os.Remove("/tmp/mocksqlite3")
	}
}

func TestNewConnection(t *testing.T) {
	sqlc := SQLConfig{
		DBName:   "demo_db",
		DBPort:   "3306",
		DBAddr:   "127.0.0.1",
		User:     "jim",
		Password: "password",
		DBType:   "mysql",
	}
	sqlc.newConnection()
	expectConection := "jim:password@tcp(127.0.0.1:3306)/demo_db?charset=utf8&parseTime=True&loc=Local"
	assert.Equal(t, expectConection, sqlc.ConnectionUrl)
}

func TestSQLWrite(t *testing.T) {
	sqlc := SQLConfig{DBType: "sqlite3"}
	sqlc.ConnectionUrl = "/tmp/mocksqlite3.db"
	err := sqlc.Write(&demo_pts)
	assert.Nil(t, err)
}

func TestOpenDB(t *testing.T) {
	mockDB, err := mockTestEnv()
	assert.Nil(t, err)
	assert.NotNil(t, mockDB)

	// opendebug mode
	mockDB.debug()

	err = mockDB.Write(&demo_pts)
	// err = mockDB.batchInsertData(&demo_pts)
	assert.Nil(t, err)
}
