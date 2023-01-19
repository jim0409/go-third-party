package main

import "log"

func main() {
	// TODO: 根據未來專案大小將設定檔移出
	mysqlAddr := "127.0.0.1"
	mysqlPort := "3306"
	mysqlOpDB := "group"
	mysqlUsr := "root"
	mysqUsrPwd := "root"

	newDB := NewDBConfiguration(mysqlUsr, mysqUsrPwd, "mysql", mysqlOpDB, mysqlPort, mysqlAddr)
	db, err := newDB.NewDBConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := db.Closed(); err != nil {
			log.Fatal(err)
		}
	}()

}
