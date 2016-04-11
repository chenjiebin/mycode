//map反射调用方法
package main

import (
	"fmt"
	//	"reflect"
)

func main() {
	var user interface{}
	user = map[string]interface{}{
		"name": "cjb",
		"age":  20,
	}
	fmt.Println(user)

	u := user.(map[string]interface{})

	val, _ := u["age"]

	//	val := reflect.ValueOf(user).MapIndex(reflect.ValueOf("age"))
	fmt.Println(val.(int))

}
