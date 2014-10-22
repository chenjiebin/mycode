package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"strings"
)

var db *sql.DB

func getAdmin(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT admin_id,admin_name FROM admin")
	checkErr(err)

	for rows.Next() {
		var adminId int
		var adminName string
		err = rows.Scan(&adminId, &adminName)
		checkErr(err)

		fmt.Fprintf(w, adminName)
	}
}

func main() {
	var err error

	host := flag.String("host", "127.0.0.1", "mysql connect host")
	port := flag.String("port", "3306", "mysql listen port")
	user := flag.String("user", "root", "mysql user")
	password := flag.String("password", "", "mysql password")

	flag.Parse()

	connectString := strings.Join([]string{*user, ":", *password, "@tcp(", *host, ":", *port, ")/app?charset=utf8"}, "")
	fmt.Println(connectString)

	db, err = sql.Open("mysql", connectString)
	checkErr(err)
	db.SetMaxIdleConns(100)

	err = db.Ping()
	checkErr(err)

	http.HandleFunc("/", getAdmin)
	err = http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
