package main

import (
	"encoding/json"
	"fmt"
)

type message struct {
	Mtype    string
	FromId   string
	FromName string
	ToId     string
	ToName   string
	Content  string
}

func main() {
	message := message{"1", "1", "cjb", "2", "iceup", "hello hehe"}
	fmt.Println(message)

	fmt.Println(json.Marshal(message))

	//student := Student{Person{2, "iceup"}, 31}
}
