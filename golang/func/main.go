package main

import (
	"fmt"
)

func main() {
	i := 2
	func() {
		fmt.Println(i)
	}()
}
