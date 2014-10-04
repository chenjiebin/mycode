package mysql

import (
	"fmt"
)

type User struct {
	Base
	table string
}

func (this *User) Find(id string) {
	fmt.Println(this.Base.Find(this.table, id))
}

var userMysql *User

func GetInstance() *User {
	return userMysql
}

func init() {
	userMysql = &User{table: "user"}
}
