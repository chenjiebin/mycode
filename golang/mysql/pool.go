//数据库持久连接和短连接、连接池性能测试
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"time"
)

var db *sql.DB

func init() {
	db, _ = sql.Open("mysql", "root:@/test?charset=utf8")
	db.SetMaxOpenConns(200)
	db.SetMaxIdleConns(100)
	db.Ping()
}

func main() {
	startHttpServer()
}

func startHttpServer() {
	http.HandleFunc("/", find)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func find(w http.ResponseWriter, r *http.Request) {
	read()
	fmt.Fprintln(w, "finish")
}

func pconnect() {
	db, _ := sql.Open("mysql", "root:@/test?charset=utf8")
	//db.Ping()
	db.SetMaxIdleConns(10)
	db.Ping()
	db.SetMaxOpenConns(30)
	//defer db.Close()

	startTime := time.Now().UnixNano()

	//ch := make(chan)

	go func() {
		var rows *sql.Rows
		var err error
		for i := 0; i < 100; i++ {
			rows, err = db.Query("SELECT * FROM user limit 1")
			checkErr(err)

			columns, _ := rows.Columns()
			scanArgs := make([]interface{}, len(columns))
			values := make([]interface{}, len(columns))
			for j := range values {
				scanArgs[j] = &values[j]
			}

			for rows.Next() {
				err = rows.Scan(scanArgs...)
				record := make(map[string]string)
				for i, col := range values {
					if col != nil {
						record[columns[i]] = string(col.([]byte))
					}
				}
				fmt.Println(record)
			}

		}
	}()

	endTime := time.Now().UnixNano()

	fmt.Println(endTime - startTime)

	time.Sleep(20 * 1e9)

}

func read() {
	for i := 0; i < 100; i++ {
		rows, err := db.Query("SELECT * FROM user limit 1")
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
	//time.Sleep(20 * 1e9)
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
