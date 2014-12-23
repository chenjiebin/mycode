package main

import (
	"fmt"
)

func hello(c chan bool) {
	fmt.Println("hello")
	c <- true
}

func main() {
	c := make(chan bool, 100)
	for j := 100; j > 0; j-- {
		//循环采集390个分类页面的详细url
		go hello(c)
	}
	for j := 100; j > 0; j-- {
		fmt.Println(<-c)
	}
}
