package goyaf

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func getConnect() *sql.DB {
	host := getConfigByKey("mysql-host")
	port := getConfigByKey("mysql-port")
	username := getConfigByKey("mysql-username")
	password := getConfigByKey("mysql-password")
	database := getConfigByKey("mysql-database")
	charset := getConfigByKey("mysql-charset")

	connectString := username + ":" + password +
		"@tcp(" + host + ":" + port + ")/" + database + "?charset=" + charset

	Debug(connectString)

	db, _ := sql.Open("mysql", connectString)

	Debug(db)

	return db
}
