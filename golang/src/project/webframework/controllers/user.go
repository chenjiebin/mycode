package controllers

import (
	"../models/business"
)

type User struct {
	Base
}

func (this *User) UserinfoAction() {
	id := this.GetRequest().GetQuery("id", "0")

	userBn := &business.User{}
	userBn.GetUserInfo(id)
}
