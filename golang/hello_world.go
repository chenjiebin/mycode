package main

import (
	//"encoding/json"
	"fmt"
	//	"math"
)

//type message struct {
//	Mtype    string
//	FromId   string
//	FromName string
//	ToId     string
//	ToName   string
//	Content  string
//}

func main() {
	array := [10]int{}
	for _, v := range array {
		fmt.Println(v)
	}
}

func add_name(i, j int) int {
	return i + j
}
