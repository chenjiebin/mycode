package mysql

import (
	"fmt"
	"goyaf"
)

type User struct {
	Base
}

func (this *User) Find(id string) {
	tableGateway := &goyaf.Table{Table: "user"}

	where := make(map[string]string)
	//where["user_id"] = id

	result := tableGateway.Select(where)

	for k, v := range result {
		fmt.Println(k, v)
	}

	fmt.Println(tableGateway.Select(where))
	//fmt.Println(this.Base.Find(this.table, id))
}

func (this *User) Add(data map[string]string) {
	tableGateway := &goyaf.Table{Table: "user"}
	tableGateway.Insert(data)
}

func (this *User) Update(data map[string]string, where map[string]string) {
	tableGateway := &goyaf.Table{Table: "user"}
	tableGateway.Update(data, where)
}

var userMysql *User

func GetUserInstance() *User {
	return userMysql
}

func init() {
	userMysql = &User{}
}
