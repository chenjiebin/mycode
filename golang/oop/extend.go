//oop extend演示
//演示改写父类的方法
//演示this方式的显式
//演示构造方法
//演示多态
package main

import (
	"fmt"
)

type Animal struct {
	Name string //动物名称
	age  int    //动物年龄
}

//声明方法
func (a Animal) eat() {
	fmt.Println(a.Name, "eat")
}

type Duck struct {
	Animal
}

func main() {
	duck := Duck{}
	duck.Animal.Name = "duck"
	duck.Animal.eat()
}
