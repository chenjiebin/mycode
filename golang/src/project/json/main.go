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

func main() {
	authors := []string{"XuShiwei", "HughLv", "Pandaman", "GuaguaSong", "HanTuo", "BertYuan", "XuDaoli"}
	gobook := Book{
		"Go语言编程",
		authors,
		"ituring.com.cn",
		true,
		9.99,
	}

	gobookJson, _ := json.Encoder(gobook)

	fmt.Println(gobookJson)

	var decodeGoBook Book

	json.Unmarshal(gobookJson, &decodeGoBook)

	fmt.Println(decodeGoBook)
}
