package mysql

import ()

type User struct {
	Base
}

var UserMysql *User

func init() {
	UserMysql = &User{}
	UserMysql.table = "user"
	UserMysql.primaryKey = "user_id"
}
