package controllers

import (
	"../models/business"
)

type User struct {
	Base
}

func (this *User) AddAction() {
	params := make(map[string]string)
	params["user_name"] = "cjb"
	params["user_age"] = "20"

	userBn := &business.User{}
	userBn.Add(params)
}

func (this *User) UpdateAction() {
	params := make(map[string]string)
	params["user_id"] = "1"
	params["user_name"] = "cjb"
	params["user_age"] = "20"

	userBn := &business.User{}
	userBn.Update(params)
}

func (this *User) UserinfoAction() {
	id := this.GetRequest().GetQuery("id", "0")

	userBn := &business.User{}
	userBn.GetUserInfo(id)
}
