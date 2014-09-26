package business

import (
	"../mysql"
)

type User struct {
}

func (this *User) GetUserInfo(id string) {
	userMysql := mysql.GetInstance()
	userMysql.Find(id)
}
