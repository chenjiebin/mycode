package main

import (
	"fmt"
	"os"
)

func main() {
	// 获取当前工作目录
	fmt.Println(os.Getwd())
	os.Chdir("/Users/chenjiebin/Sites/mycode/golang")
	fmt.Println(os.Getwd())
}
