package main

import (
	"fmt"
)

type User struct {
	Id   int
	Name string
}

func (this *User) SetName() {
	this.Name = "hehe"
}

func main() {
	users := map[string]User{
		"1": User{Id: 1, Name: "cjb"},
	}

	user := users["1"]

	user2 := &user

	fmt.Println(user2)
	user2.Id = 2
	fmt.Println(user2)
	user2.SetName()
	fmt.Println(user2)

	fmt.Println(users)
}
