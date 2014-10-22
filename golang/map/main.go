package main

import (
	"fmt"
)

func main() {
	names := make(map[int]string)
	names[1] = "cjb"

	for k, v := range names {
		fmt.Println(k, "=", v)
	}

	fmt.Println(names)
}
