package mysqldb

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMySqlDBAccessObj(t *testing.T) {
	ins1 := RetriveMySQLDBAccessObj()
	ins2 := RetriveMySQLDBAccessObj()
	assert.Equal(t, ins1, ins2)
}

func TestMySQLDB(t *testing.T) {

}

func TestNewConnection(t *testing.T) {
	os.Setenv("TestDB", "test")
	LoadMySQLDBConfig("mysql", "127.0.0.1", "3306", "root", "secret", true, 10, 10)
	StartMySQLDB()

	migrationA()

	ins := RetriveMySQLDBAccessObj()
	assert.NoError(t, ins.CreateRecord("jim", "CN", "EN"))

	record, err := ins.QueryRecord("jim")
	assert.NoError(t, err)
	// assert.Equal(t, "", *record)
	for _, j := range *record {
		for _, jj := range j.Users {
			assert.Equal(t, "jim", jj.Name)
		}
	}

}
