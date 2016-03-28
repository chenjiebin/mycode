package main

import (
	"fmt"
)

func main() {
	a := []byte{0, 1, 1, 1, 1}
	b := []byte{0, 1, 1, 1, 0}

	for k, v := range a {
		fmt.Println(v | b[k])
	}
}
