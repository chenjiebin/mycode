//结构体反射调用方法
package main

import (
	"fmt"
	"reflect"
)

type Base struct {
}

func (this *Base) Hello() {
	fmt.Println("hello")
}

type User struct {
	Base
	id   int
	name string
}

func (this *User) SayHello() string {
	return "Hello"
}

func (this *User) SayWord(word string) {
	fmt.Println(word)
}

func main() {
	var user interface{}
	user = &User{id: 22, name: "iceup"}
	fmt.Println(user) //就是检查一下myType对象内容

	reflect.ValueOf(user).MethodByName("Hello").Call(nil)

	val := reflect.ValueOf(user).MethodByName("SayHello").Call(nil)[0]
	fmt.Println(val)

	params := make([]reflect.Value, 1)
	params[0] = reflect.ValueOf("Hehe")
	reflect.ValueOf(user).MethodByName("SayWord").Call(params)

	//for i := 0; i < val.NumField(); i++ {
	//	valueField := val.Field(i)
	//	typeField := val.Type().Field(i)
	//	tag := typeField.Tag

	//	fmt.Printf("Field Name: %s,\t Field Value: %v,\t Tag Value: %s\n", typeField.Name, valueField.Interface(), tag.Get("tag_name"))
	//}

}
