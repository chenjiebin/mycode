package dao

import (
	"../mysql"
)

type User struct {
	Base
}

var UserDAO *User

func init() {
	UserDAO = &User{}
	UserDAO.mysql = mysql.UserMysql
}
