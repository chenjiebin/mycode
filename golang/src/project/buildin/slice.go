package main

import (
	"fmt"
)

func main() {
	var myArray [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(myArray[5:10])
	fmt.Println(myArray[:5])
}
