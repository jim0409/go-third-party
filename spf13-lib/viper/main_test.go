package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	logRoot     = "service.log"
	mysqlRoot   = "service.mysql"
	serviceRoot = "service.service"
)

func TestLoadConfig(t *testing.T) {
	configPath := "./config.yaml"

	Config := &Configs{}
	Config.loadConfig(configPath)

	assert.NotEqual(t, Config, "")
	assert.Equal(t, Config.GetLogLevel(), "info")
	assert.Equal(t, Config.GetMysqlTable1(), "table1")
	assert.Equal(t, Config.GetMysqlTable2(), "")
}

// GetLogLevel : Level

func (c *Configs) GetLogLevel() string {
	return c.config.GetString(logRoot + ".level")
}

// Get Mysql User Info
func (c *Configs) GetMysqlHost() string {
	return c.config.GetString(mysqlRoot + ".host")
}

func (c *Configs) GetMysqlPort() string {
	return c.config.GetString(mysqlRoot + ".port")
}

func (c *Configs) GetMysqlUser() string {
	return c.config.GetString(mysqlRoot + ".user")
}

func (c *Configs) GetMysqlPasswd() string {
	return c.config.GetString(mysqlRoot + ".passwd")
}

func (c *Configs) GetMysqlMaxConn() int {
	return c.config.GetInt(mysqlRoot + ".max_conn")
}

func (c *Configs) GetName() string {
	return c.config.GetString(serviceRoot + ".name")
}

func (c *Configs) GetPort() string {
	return c.config.GetString(serviceRoot + ".port")
}

// Get Mysql Tables : Table1/ Table2/ Table3

func (c *Configs) GetMysqlTable1() string {
	return c.config.GetString(mysqlRoot + ".database.table1")
}

func (c *Configs) GetMysqlTable2() string {
	return c.config.GetString(mysqlRoot + ".database.table2")
}

func (c *Configs) GetMysqlTable3() string {
	return c.config.GetString(mysqlRoot + ".database.table3")
}
