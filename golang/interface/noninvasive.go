//演示非侵入式的方法
//在这里animal为一个接口类型，声明了eat和play方法
//任何对象只要实现了这两个方法，就可以称为实现了该接口
//这种实现并非是强制性的，当然如果没有完全实现对应接口的方法，亦不能作为该接口的类型
//例如dog对象如果没有实现play方法，则当参数传递给see函数，执行是会报错的。
package main

import (
	"fmt"
)

type animal interface {
	eat()
	play()
}

type dog struct {
}

func (this *dog) eat() {
	fmt.Println("dog eat")
}

func (this *dog) play() {
	fmt.Println("dog play")
}

func see(a animal) {
	a.eat()
}

func main() {
	d := &dog{}
	see(d)
}
