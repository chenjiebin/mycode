package main

import (
	"fmt"
)

func main() {
	buffer := make([]byte, 1000)
	//fmt.Println(buffer[0:1000])

	n := len(buffer)

	for i := 0; i < n; i++ {
		fmt.Println("i=", i, "i+20=", i+20)
		if n < i+20 {
			break
		}

		fmt.Println(buffer[i : i+20])
	}

	fmt.Println(len(buffer))
}
