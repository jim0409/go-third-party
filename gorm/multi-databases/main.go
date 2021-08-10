package main

import (
	"go-third-party/gorm/multi-databases/models"
)

var (
	usr    = "jim"
	pwd    = "password"
	dbt    = "mysql"
	dbname = "db"
	port   = "3306"
	addr   = "127.0.0.1"
)

func main() {
	m := models.InitMainDB(usr, pwd, dbt, dbname, port, addr)
	if err := m.CreateMessage("msg-1", "jim1", "31"); err != nil {
		panic(err)
	}
	if err := m.CreateMessage("msg-2", "jim2", "32"); err != nil {
		panic(err)
	}
	if err := m.CreateMessage("msg-0", "jim0", "30"); err != nil {
		panic(err)
	}
}
