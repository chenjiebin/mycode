package main

import (
	"fmt"
	"reflect"
)

func main() {
	//float64Test()
	//valueOfTest()
	//structTest()
	//canSet()
	structTest()
}

func float64Test() {
	var x float64 = 3.4
	fmt.Println("type:", reflect.TypeOf(x))
}

func valueOfTest() {
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())
	fmt.Println("kind is ", v.Kind())
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("value:", v.Float())
}

func canSet() {
	var x float64 = 3.4
	p := reflect.ValueOf(&x) // 注意:得到X的地址
	fmt.Println("type of p:", p.Type())
	fmt.Println("settability of p:", p.CanSet())
	v := p.Elem()
	fmt.Println("settability of v:", v.CanSet())
	v.SetFloat(7.1)
	fmt.Println(v.Interface())
	fmt.Println(x)
}

func structTest() {
	type T struct {
		A int
		B string
	}

	t := T{203, "mh203"}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
}
