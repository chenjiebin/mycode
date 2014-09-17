//结构体反射调用方法
package main

import (
	"fmt"
	"reflect"
)

type User struct {
	id   int
	name string
}

func (this *User) SayHello() string {
	return "Hello"
}

func main() {
	var user interface{}
	user = &User{22, "iceup"}
	fmt.Println(user) //就是检查一下myType对象内容

	val := reflect.ValueOf(user).MethodByName("SayHello").Call(nil)[0]

	fmt.Println(val)

	//for i := 0; i < val.NumField(); i++ {
	//	valueField := val.Field(i)
	//	typeField := val.Type().Field(i)
	//	tag := typeField.Tag

	//	fmt.Printf("Field Name: %s,\t Field Value: %v,\t Tag Value: %s\n", typeField.Name, valueField.Interface(), tag.Get("tag_name"))
	//}

}
