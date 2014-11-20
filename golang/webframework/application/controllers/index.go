package controllers

import (
	"fmt"
	"git.oschina.net/iceup/goyaf/db"
)

type Index struct {
	Base
}

func (this *Index) IndexAction() {
	var body string
	body = "index index index"

	id := this.GetRequest().GetParam("id", "0")
	body = body + "<br />params id:" + id

	this.GetResponse().AppendBody(body)

}

//数据库操作demo
func (this *Index) DbdemoAction() {
	config := map[string]string{
		"driver":   "mysql",
		"host":     "127.0.0.1",
		"port":     "3306",
		"database": "test",
		"username": "root",
		"password": "",
		"charset":  "utf8",
	}

	test["mysql-host"] = "192.168.3.233"
	test["mysql-port"] = "3306"
	test["mysql-database"] = "pba-user"
	test["mysql-username"] = "root"
	test["mysql-password"] = "123456"
	test["mysql-charset"] = "utf8"
	test["mysql-driver_options-1002"] = "SET NAMES utf8mb4"

	adapter := db.NewAdapter(config)
	userTable := db.NewTable("user", adapter)

	//查询数据
	result, err := userTable.Select(db.Select{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)

	//插入数据
	lastInsertId, err := userTable.Insert(map[string]string{"user_name": "iceup", "user_age": "20"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(lastInsertId)

	//更新数据
	affect, err := userTable.Update(map[string]string{"user_name": "cjb", "user_age": "21"}, map[string]string{"user_id": "1"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(affect)

	//删除数据
	delaffect, err := userTable.Delete(map[string]string{"user_id": "1"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(delaffect)
}

func init() {
	fmt.Println("init controller index")
}
