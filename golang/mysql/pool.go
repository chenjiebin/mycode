//数据库持久连接和短连接、连接池性能测试
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func main() {
	pconnect()
}

func pconnect() {
	db, _ := sql.Open("mysql", "root:@/test?charset=utf8")
	//db.Ping()
	db.SetMaxIdleConns(100)
	db.Ping()
	//db.SetMaxOpenConns(1)
	//defer db.Close()

	startTime := time.Now().UnixNano()

	var rows *sql.Rows
	var err error
	for i := 0; i < 100; i++ {
		rows, err = db.Query("SELECT * FROM user limit 1")
		checkErr(err)
		db.Close()
	}

	endTime := time.Now().UnixNano()

	fmt.Println(endTime - startTime)

	fmt.Println(rows)

	time.Sleep(30 * 1e9)

}

func connect() {
	db, err := sql.Open("mysql", "89test:89test@tcp(192.168.3.89:3306)/business?charset=utf8")
	checkErr(err)

	rows, err := db.Query("SELECT * FROM tb_users limit 1")
	checkErr(err)

	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		//将行数据保存到record字典
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

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
