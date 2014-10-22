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

func (this *User) Add(params map[string]string) {
	userMysql := mysql.GetUserInstance()
	userMysql.Add(params)
}

func (this *User) Update(params map[string]string) {
	data := make(map[string]string)
	data["user_name"] = params["user_name"]
	data["user_age"] = params["user_age"]

	where := make(map[string]string)
	where["user_id"] = params["user_id"]

	userMysql := mysql.GetUserInstance()
	userMysql.Update(data, where)
}
