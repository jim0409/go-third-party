package models

import (
	conf "go-third-party/gin/customerized/conf"
	mysqldb "go-third-party/gin/customerized/models/mysqldb"
	redisdb "go-third-party/gin/customerized/models/redisdb"
)

/*
use facade design pattern to new db object
1. initDB
2. closeDB
*/

func RetriveRedisAccessModel() RedisImplement {
	return redisdb.RetriveRedisPoolObj()
}

type RedisImplement interface {
	GetKeyValue(string) (interface{}, error)
	SetKey(string, interface{}) error
	SetExpiredKey(string, interface{}, int) error
	SetNonExistedKey(string, interface{}) error
	DeleteKey(string) error
	Close() error
}

func RetriveMySqlDbAccessModel() MySqlImplement {
	return mysqldb.RetriveMySQLDBAccessObj()
}

type MySqlImplement interface {
	ExecSql(string) error
	Close() error
}

// InitDb init db
func InitDb(mysqlconf *conf.DbConf, redisconf *conf.RedisConf) error {
	var err error
	if err = initMySql(mysqlconf); err != nil {
		return err
	}
	if err = initRedis(redisconf); err != nil {
		return err
	}

	return nil
}

// Close close db
func Close() error {
	RetriveMySqlDbAccessModel().Close()
	RetriveRedisAccessModel().Close()
	return nil
}

func initMySql(c *conf.DbConf) error {
	mysqldb.LoadMySQLDBConfig(
		c.DbName,
		c.DbHost,
		c.DbPort,
		c.DbUser,
		c.DbPassword,
		c.DbLogEnable,
		c.DbMaxConnect,
		c.DbIdleConnect,
	)
	return mysqldb.StartMySQLDB()
}

func initRedis(c *conf.RedisConf) error {
	redisdb.LoadRedisDBConfig(
		c.RedisAddr,
		c.RedisAuth,
		c.RedisMaxIdle,
		c.RedisMaxActive,
		c.RedisIdleTimeout,
	)
	return redisdb.StartRedisPool()
}
