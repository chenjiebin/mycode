package main

import (
	"fmt"
	"os"
)

func main() {
	pid := os.Getppid()

	fmt.Println(pid)
}
