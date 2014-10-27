package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

/*获取当前文件执行的路径*/
func GetCurrPath() string {
	file, _ := exec.LookPath(os.Args[0])
	fmt.Println(file)

	path, _ := filepath.Abs(file)
	fmt.Println(path)
	splitstring := strings.Split(path, "\\")
	size := len(splitstring)
	splitstring = strings.Split(path, splitstring[size-1])
	ret := strings.Replace(splitstring[0], "\\", "/", size-1)
	return ret
}

func main() {
	fmt.Println(GetCurrPath())

}
