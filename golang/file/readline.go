// 一行行读取文件内容
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("./file.txt")
	if err != nil {
		panic(err)
	}
	buff := bufio.NewReader(f)
	for i := 1; ; i++ {
		line, err := buff.ReadBytes('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}

		fmt.Printf("%d line: %s", i, string(line))

		// 文件已经到达结尾
		if err == io.EOF {
			break
		}
	}

	fmt.Println()
}
