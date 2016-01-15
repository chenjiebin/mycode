package main

import (
	"fmt"
	"os"
)

func main() {
	//声明变量
	var name string
	//赋值
	name = "meiqu"
	//声明并赋值
	//	name := "meiqu"

	//循环条件没有小刮号
	for i := 1; i < 10; i++ {

	}
}

func demo() (name string, err error) {
switch i { case 0:
fmt.Printf("0") case 1:
fmt.Printf("1") case 2:
fallthrough case 3:
fmt.Printf("3") case 4, 5, 6:
fmt.Printf("4, 5, 6") default:
            fmt.Printf("Default")
}
