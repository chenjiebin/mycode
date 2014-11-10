package main

import (
	"fmt"
	"reflect"
)

type Base struct {
	Age int
}

func (this *Base) SetAge() {
	this.Age = 90
}

type User struct {
	Base
	Id   int
	Name string
}

func (this *User) SetName() {
	this.Name = "iceup"
}

func main() {
	var u interface{}
	u = User{Id: 1, Name: "cjb"}
	fmt.Println(u)

	//fmt.Println("u field:", reflect.ValueOf(u).NumField())

	//fmt.Println(reflect.ValueOf(u).Field(0), reflect.ValueOf(u).Field(1), reflect.ValueOf(u).Field(2))

	t := reflect.ValueOf(u).Type()
	fmt.Println(t)
	v := reflect.New(t)
	fmt.Println(v)
	//fmt.Println(v.NumField())
	fmt.Println(v.MethodByName("SetAge").Call(nil))
	fmt.Println(v.Interface())

	//fmt.Println(reflect.ValueOf(&u2).MethodByName("SetName").Call(nil))
	//fmt.Println(reflect.ValueOf(u2).MethodByName("SetAge").Call(nil))

}
