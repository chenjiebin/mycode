package main

import (
	"fmt"
	"reflect"
)

type base struct {
	age int
}

//type controller interface {
//	Hello()
//}

func (this *base) Hello() {
	fmt.Println("hello")
	this.age = 10

}

type user struct {
	base
	Id   int64
	name string
}

func (this user) To() {
	this.name = "hehe"
}

func main() {
	var u interface{}
	u = user{Id: 1, name: "cjb"}
	fmt.Println(reflect.ValueOf(&u).Elem())
	fmt.Println(u)

	//u2 := u
	//u3 := &u2

	//fmt.Println(c)
	//fmt.Println(u2)

}
