package mysqldb

// rewrite it to registry table struct to registry.go
func migrationA() {
	r := RetriveMySQLDBAccessObj()
	r.migration(&User{})
	r.migration(&Language{})
}
