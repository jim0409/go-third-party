package dbpkg

import (
	"context"
	"log"
	"net"
	"os"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

// func init() {
// 	clearTestEnv()
// }

func clearTestEnv() {
	log.Println("Clear test env")
	if _, err := os.Stat("/tmp/gorm.db"); !os.IsNotExist(err) {
		log.Println("Remove Origin gorm.db Files")
		os.Remove("/tmp/gorm.db")
	}
}

// Set mock DB env
func mockTestEnv() (OPDB, error) {
	var dbc = DBConfig{}
	dbc.DBUri = "/tmp/gorm.db"
	dbc.DBType = "sqlite3"
	opdb, err := dbc.NewDBConnection()
	return opdb, err
}

func TestNewConnection(t *testing.T) {
	var (
		user     = "jim"
		password = "pw"
		dbtype   = "sqlite"
		dbname   = "demo_db"
		dbport   = "3306"
		dbaddr   = "127.0.0.1"
	)

	NewDBConfig := NewDBConfiguration(user, password, dbtype, dbname, dbport, dbaddr)
	assert.Equal(t, user, NewDBConfig.User)
	assert.Equal(t, password, NewDBConfig.Password)
	assert.Equal(t, dbtype, NewDBConfig.DBType)
	assert.Equal(t, dbname, NewDBConfig.DBName)
	assert.Equal(t, dbport, NewDBConfig.DBPort)
	assert.Equal(t, dbaddr, NewDBConfig.DBAddress)
	assert.Equal(t, "jim:pw@tcp(127.0.0.1:3306)/demo_db?charset=utf8&parseTime=True&loc=Local", NewDBConfig.DBUri)

}

func TestNewDBOperation(t *testing.T) {
	opdb, err := mockTestEnv()
	assert.Nil(t, err)

	dt := &DemoTable{
		Name:  "jim",
		Email: "email@example.com",
	}
	update_email := "update_email@example.com"

	opdb.Debug()

	err = opdb.Create(dt.Name, dt.Email)
	assert.Nil(t, err)

	resStr, err := opdb.QueryWithName(dt.Name)
	assert.Equal(t, dt.Email, resStr)
	assert.Nil(t, err)

	err = opdb.UpdateEmail(dt.Name, update_email)
	assert.Nil(t, err)

	resStr, err = opdb.QueryWithName(dt.Name)
	assert.Equal(t, update_email, resStr)

	err = opdb.DeleteData(dt.Name, update_email)
	assert.Nil(t, err)

	resStr, err = opdb.QueryWithName(dt.Name)
	assert.Equal(t, "Can't find the email with jim", resStr)
	assert.Error(t, err)

}

func BenchmarkCreate(b *testing.B) {

	mysqlAddr := "127.0.0.1"
	mysqlPort := "3306"
	mysqlOpDB := "testdb"
	mysqlUsr := "jim"
	mysqUsrPwd := "password"

	newDB := NewDBConfiguration(mysqlUsr, mysqUsrPwd, "mysql", mysqlOpDB, mysqlPort, mysqlAddr)
	db, err := newDB.NewDBConnection()
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		panic(err)
	}

	dt := &DemoTable{
		Name:  "jim",
		Email: "email@example.com",
	}

	for i := 0; i < b.N; i++ {
		if err := db.Create(dt.Name, dt.Email); err != nil {
			panic(err)
		}
	}

}

// 配合 ../change_dns_resolver
func TestForDnsCheck(t *testing.T) {
	net.DefaultResolver = &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{
				// Timeout: time.Millisecond * time.Duration(10000),
				Timeout: time.Second * time.Duration(10),
			}
			return d.DialContext(ctx, network, "127.0.0.1:5301")
		},
	}

	mysqlAddr := "dns.jim.host"
	mysqlPort := "3306"
	mysqlOpDB := "testdb"
	mysqlUsr := "jim"
	mysqUsrPwd := "password"

	newDB := NewDBConfiguration(mysqlUsr, mysqUsrPwd, "mysql", mysqlOpDB, mysqlPort, mysqlAddr)
	db, err := newDB.NewDBConnection()
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		panic(err)
	}

	dt := &DemoTable{
		Name:  "jim",
		Email: "email@example.com",
	}

	for i := 0; i < 20000; i++ {
		if err := db.Create(dt.Name, dt.Email); err != nil {
			panic(err)
		}
	}
}

// go test -benchmem -run=^$ -bench ^BenchmarkForDnsCheck$ -count 1 -v
func BenchmarkForDnsCheck(b *testing.B) {
	net.DefaultResolver = &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{
				// Timeout: time.Millisecond * time.Duration(10000),
				Timeout: time.Second * time.Duration(10),
			}
			return d.DialContext(ctx, network, "127.0.0.1:5301")
		},
	}

	mysqlAddr := "dns.jim.host"
	mysqlPort := "3306"
	mysqlOpDB := "testdb"
	mysqlUsr := "jim"
	mysqUsrPwd := "password"

	newDB := NewDBConfiguration(mysqlUsr, mysqUsrPwd, "mysql", mysqlOpDB, mysqlPort, mysqlAddr)
	db, err := newDB.NewDBConnection()
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		panic(err)
	}

	dt := &DemoTable{
		Name:  "jim",
		Email: "email@example.com",
	}

	for i := 0; i < b.N; i++ {
		if err := db.Create(dt.Name, dt.Email); err != nil {
			panic(err)
		}
	}
}
