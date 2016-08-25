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

	fmt.Println(arrs[5:6])
	fmt.Println(arrs[9:10])
}
