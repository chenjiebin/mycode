//oop struct演示
//演示实例化对象的加&的方式和不加&的方法
package main

import (
	"fmt"
)

//声明动物这个struct
//类似
//class Animal {
//	public String name
//	private int age
//}
//
type Animal struct {
	Name string //动物名称
	age  int    //动物年龄
}

//声明方法
func (a Animal) eat() {
	fmt.Println(a.Name, "eat")
}

func main() {
	animal := Animal{Name: "duck"}
	animal.age = 10
	fmt.Println(animal)

	animal.eat()
}
