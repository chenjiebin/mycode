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

func (this *User) Add(params map[string]string) {
	tableGateway := &goyaf.Table{Table: "user"}
	tableGateway.Insert(params)
}

var userMysql *User

func GetUserInstance() *User {
	return userMysql
}

func init() {
	userMysql = &User{}
}
