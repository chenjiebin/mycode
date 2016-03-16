// linux上通过offset写文件的时候会原先该位置的内容往后偏移
// 而后在该位置上写入内容
package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	offs := flag.Int64("offs", 0, "offs")
	flag.Parse()

	fmt.Println(*offs)

	os.op
	file, err := os.OpenFile("hole", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	defer file.Close()

	b := []byte{1}
	n, err := file.WriteAt(b, *offs)

	fmt.Println(n, err)
	// output: 1 <nil>
}
