package main

import (
	"fmt"
	"go-third-party/gorm/multi-databases/models"
	"testing"
)

func TestWriteMsg(t *testing.T) {
	m := models.InitMainDB(usr, pwd, dbt, dbname, port, addr)
	defer func() {
		if err := m.Closed(); err != nil {
			panic(err)
		}
	}()

	for i := 0; i < 10; i++ {
		id, g, n, a := rGen()
		fmt.Println(g, n, a)
		// .. m.CreateMessage("msg-0", "jim0", "30") ..
		if err := m.CreateMessage(g, n, a); err != nil {
			panic(err)
		}
		fmt.Println(id)
	}
}
