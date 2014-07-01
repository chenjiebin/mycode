/*
	处理过期的积分兑换商品，返回用户积分
*/
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@tcp(192.168.3.233:3306)/app?charset=utf8")
	checkErr(err)

	rows, err := db.Query("SELECT admin_id,admin_name FROM admin")
	checkErr(err)
	for rows.Next() {
		var adminId int
		var adminName string
		err = rows.Scan(&adminId, &adminName)
		checkErr(err)

		fmt.Println(adminId)
		fmt.Println(adminName)
	}

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
