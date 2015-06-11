package main

import (
	"fmt"
)

type Integer int

func (a Integer) Less(b Integer) bool {
	return a < b
}
func (a *Integer) Add(b Integer) {
	*a += b
}

type LessAdder interface {
	Less(b Integer) bool
	Add(b Integer)
}

func main() {
	var a Integer = 1
	//a.Add(1)
	//fmt.Println(a.Less(2))
	var b1 LessAdder = &a //OK
	fmt.Println(b1)
	var b2 LessAdder = a //not OK
	fmt.Println(b2)
}
