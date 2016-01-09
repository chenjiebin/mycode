package main

import "fmt"

func main() {
	c := make(chan int, 10)
	for i := 0; i < 10; i++ {
		c <- i
	}
	l := len(c)
	for i := 0; i < l; i++ {
		fmt.Println(<-c)
	}
}
