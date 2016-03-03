// 一次性读取文件的所有内容

package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("./file.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
