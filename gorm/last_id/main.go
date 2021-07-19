package main

import (
	"log"
)

func main() {
	// if len(os.Args) != 6 {
	// 	log.Fprintf(os.Stderr, "Usage: %s <msyql-addr> <mysql-port> <mysql-operation-database> <mysql-user> <mysql-user-password>\n",
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
	mysqlOpDB := "testdb"
	mysqlUsr := "jim"
	mysqUsrPwd := "password"

	newDB := NewDBConfiguration(mysqlUsr, mysqUsrPwd, "mysql", mysqlOpDB, mysqlPort, mysqlAddr)
	db, err := newDB.NewDBConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		// if err := db.cleanAll(); err != nil {
		// 	log.Printf("Error happend while cleaning %s\n", err)
		// } else {
		// 	log.Println("success drop mysql talbe demo_tables")
		// }
		if err := db.Closed(); err != nil {
			log.Fatal(err)
		}
	}()

	u := &Usr{Name: "jim"}
	if err := db.InsertUsr(u); err != nil {
		log.Println(err)
	}
	log.Println(u.ID)
}
