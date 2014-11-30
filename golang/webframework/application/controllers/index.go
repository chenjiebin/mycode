package controllers

import (
	"fmt"
	"git.oschina.net/iceup/goyaf/db"
	"strconv"
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
		"driver":      "mysql",
		"host":        "192.168.1.102",
		"port":        "3306",
		"database":    "test",
		"username":    "root",
		"password":    "",
		"charset":     "utf8",
		"maxconn":     "300",
		"maxidleconn": "100",
	}

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
	affect, err := userTable.Update(map[string]string{"user_name": "cjb", "user_age": "21"},
		map[string]string{"user_id": strconv.FormatInt(lastInsertId, 10)})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(affect)

	//删除数据
	delaffect, err := userTable.Delete(map[string]string{"user_id": strconv.FormatInt(lastInsertId, 10)})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(delaffect)
}

func init() {
	fmt.Println("init controller index")
}
