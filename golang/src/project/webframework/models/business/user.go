package business

import (
	"../mysql"
)

type User struct {
}

func (this *User) GetUserInfo(id string) {
	userMysql := mysql.GetUserInstance()
	userMysql.Find(id)
}

func (this *User) Add(data map[string]string) {
	userMysql := mysql.GetUserInstance()
	userMysql.Add(data)
}
