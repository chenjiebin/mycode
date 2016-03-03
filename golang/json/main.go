package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title       string
	Authors     []string
	Publisher   string
	IsPublished bool
	Price       float64
}

type User struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Message string `json:"message"`
}

func main() {
	userJsonStr, _ := json.Marshal(User{Id: 1, Name: "golang", Message: "message"})
	fmt.Println(string(userJsonStr))

	//authors := []string{"XuShiwei", "HughLv", "Pandaman", "GuaguaSong", "HanTuo", "BertYuan", "XuDaoli"}
	//gobook := Book{
	//	Title: "Go语言编程",
	//	authors,
	//	"ituring.com.cn",
	//	true,
	//	9.99,
	//}

	//gobookJson, _ := json.Marshal(gobook)

	//fmt.Println(gobookJson)

	//var decodeGoBook Book

	//json.Unmarshal(gobookJson, &decodeGoBook)

	//fmt.Println(decodeGoBook)
}
