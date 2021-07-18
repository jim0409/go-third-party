package main

var (
	usr    = "jim"
	pwd    = "password"
	dbt    = "mysql"
	dbname = "main"
	port   = "3306"
	addr   = "127.0.0.1"
)

func main() {
	m := InitMainDB(usr, pwd, dbt, dbname, port, addr)
	m.StartShardDbs()
	if err := m.CreateMessage("msg1", "jim", "32"); err != nil {
		panic(err)
	}
}
