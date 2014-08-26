package main

import (
	"fmt"
	//"reflect"
)

type Human struct {
	Name string
	Age  int
}
type Man struct {
	Name string
	Age  int
}

type Woman struct {
	Human
}

func main() {
	var obj interface{}
	obj = Man{Age: 20}

	var man Man

	man = obj.(Man)

	fmt.Println(man.Age)

	//obj.Name = "cjb"
	//obj.Age = 10
	//fmt.Println(obj)
}
