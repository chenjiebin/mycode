// golang创建目录
package main

import (
	"fmt"
	"os"
)

func main() {
	isDirExist()
}

// 递归创建目录
func makedir() {
	err := os.Mkdir("tmp2", 0777)
	fmt.Println(err)
}

// 递归创建目录
func makedirAll() {
	err := os.MkdirAll("tmp/name1/name2", 0777)
	fmt.Println(err)
}

// 判断目录是否存在
func isDirExist() {
	path := "tmp3"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Println(err)
	}
}
