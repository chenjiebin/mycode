package main

import (
	//"encoding/json"
	"fmt"
	"math"
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
	total := 0.0
	for i := 0; i < 240; i++ {
		total += 1700 * math.Pow(1+0.041910/360, float64(7200-i*30))
	}
	fmt.Println(total)
	//1700 * (1 + 0.05/360)
	//message := message{"1", "1", "cjb", "2", "iceup", "hello hehe"}
	//fmt.Println(message)

	//fmt.Println(json.Marshal(message))

	//student := Student{Person{2, "iceup"}, 31}
}
