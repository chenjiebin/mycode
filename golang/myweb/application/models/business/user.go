package business

import (
	"../dao"
	"fmt"
)

type User struct {
}

func (this *User) GetUserInfo(id string) {
	dao.UserDAO.Find(id)
}

func (this *User) Add(params map[string]string) {
	dao.UserDAO.Insert(params)
}

func (this *User) Update(params map[string]string) {
	data := make(map[string]string)
	data["user_name"] = params["user_name"]
	data["user_age"] = params["user_age"]

	where := make(map[string]string)
	where["user_id"] = params["user_id"]

	dao.UserDAO.Update(data, where)
}

var UserBN *User

func init() {
	fmt.Println("init business user")
	UserBN = &User{}
}
