package main

import (
	"fmt"
)

func main() {
	arrs := make([]int, 10)
	for i := 0; i < 10; i++ {
		arrs[i] = i
	}
	fmt.Println(arrs)
}
