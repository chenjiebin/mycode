package mysql

import (
	"fmt"
)

type User struct {
	Base
}

func (this *User) Find(id string) {
	fmt.Println(this.Base.Find(id))
	//fmt.Println(this.Base.Find(this.table, id))
}

var userMysql *User

func GetUserInstance() *User {
	return userMysql
}

func init() {
	userMysql = &User{}
	userMysql.Table = "user"
}
