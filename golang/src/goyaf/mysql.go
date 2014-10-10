package goyaf

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func getConnect() *sql.DB {
	host := GetConfigByKey("mysql-host")
	port := GetConfigByKey("mysql-port")
	username := GetConfigByKey("mysql-username")
	password := GetConfigByKey("mysql-password")
	database := GetConfigByKey("mysql-database")
	charset := GetConfigByKey("mysql-charset")

	connectString := username + ":" + password +
		"@tcp(" + host + ":" + port + ")/" + database + "?charset=" + charset

	Debug(connectString)

	db, _ := sql.Open("mysql", connectString)

	Debug(db)

	return db
}
