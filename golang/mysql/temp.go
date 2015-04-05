package main

import (
	_ "./go-sql-driver/mysql"
	"database/sql"
	"fmt"
)

func main() {
	query()
}

//查询demo
func query() {
	db, err := sql.Open("mysql", "root:@/test?charset=utf8")
	checkErr(err)

	rows, err := db.Query("SELECT * FROM user")
	checkErr(err)

	//普通demo
	//for rows.Next() {
	//	var userId int
	//	var userName string
	//	var userAge int
	//	var userSex int

	//	rows.Columns()
	//	err = rows.Scan(&userId, &userName, &userAge, &userSex)
	//	checkErr(err)

	//	fmt.Println(userId)
	//	fmt.Println(userName)
	//	fmt.Println(userAge)
	//	fmt.Println(userSex)
	//}

	//字典类型
	//构造scanArgs、values两个数组，scanArgs的每个值指向values相应值的地址
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
		panic(err)
	}
}
