package db

import (
	"database/sql"
	"fmt"
)

var DB *sql.DB

func init() {
	DB, _ = sql.Open("mysql", "root:@/test?charset=utf8")
	fmt.Println(DB)
}
