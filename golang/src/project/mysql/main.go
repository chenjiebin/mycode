package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	query()
}
func insert() {
	db, err := sql.Open("mysql", "root:@/test?charset=utf8")
	checkErr(err)
	//插入数据
	stmt, err := db.Prepare("INSERT user SET name=?")
	checkErr(err)
	res, err := stmt.Exec("cjb")
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)
}

func query() {
	db, err := sql.Open("mysql", "root:@/test?charset=utf8")
	checkErr(err)
	//查询数据

	rows, err := db.Query("SELECT * FROM user")
	checkErr(err)

	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	//fmt.Println(user)

	//users = make(map[string]map[string]string)

	for rows.Next() {
		err = rows.Scan(scanArgs...)

		fmt.Println(values)

		record := make(map[string]string)

		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}

		fmt.Println(record)
	}

	//for rows.Next() {
	//	var uid int
	//	var username string
	//	var department string
	//	var created string

	//	rows.Columns()
	//	err = rows.Scan(&uid, &username, &department, &created)
	//	checkErr(err)
	//	fmt.Println(uid)
	//	fmt.Println(username)
	//	fmt.Println(department)
	//	fmt.Println(created)
	//}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
