//数据库持久连接和短连接、连接池性能测试
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

var db *sql.DB

func init() {
	db, _ = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test?charset=utf8")
	db.SetMaxOpenConns(2000)
	db.SetMaxIdleConns(1000)
	db.Ping()
}

func main() {
	startHttpServer()
}

func startHttpServer() {
	http.HandleFunc("/pool", pool)
	http.HandleFunc("/nopool", nopool)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func pool(w http.ResponseWriter, r *http.Request) {
	queryFromUser(db)
	fmt.Fprintln(w, "finish")
}

func nopool(w http.ResponseWriter, r *http.Request) {
	db2, _ := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test?charset=utf8")
	defer db2.Close()

	queryFromUser(db2)

	fmt.Fprintln(w, "finish")
}

func queryFromUser(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM user limit 1")
	defer rows.Close()
	checkErr(err)

	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for j := range values {
		scanArgs[j] = &values[j]
	}

	record := make(map[string]string)
	for rows.Next() {
		//将行数据保存到record字典
		err = rows.Scan(scanArgs...)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
	}

	fmt.Println(record)
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
