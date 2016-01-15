//oop类型演示
//这种方式可以很方便对内建类型进行扩展
//演示自增方法中传递指针和非指针的区别
package main

import (
	"fmt"
)

//定义一个Integer类型
type Integer int

//定义Integer类型的一个方法
func (a Integer) Less(b Integer) bool {
	return a < b
}

//主函数
func main() {
	var a Integer = 1
	fmt.Println(a.Less(2))
}
