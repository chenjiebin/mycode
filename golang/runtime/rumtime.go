package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU() - 1)
	fmt.Println(runtime.NumCPU())
}
