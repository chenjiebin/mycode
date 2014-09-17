package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(len(strings.Split("/index/index/", "/")))
}
