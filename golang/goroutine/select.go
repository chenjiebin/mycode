package main

import (
	"fmt"
	"time"
)

func foo(c chan int) {
	for {
		select {
		case i := <-c:
			fmt.Println("foo", i)
		default:
			fmt.Println("foo default")
			break

		}
	}
	fmt.Println("foo exit")
}

func main() {
	c := make(chan int)
	go foo(c)
	//	for i := 0; i < 100; i++ {
	//		c <- i
	//	}

	for {
		time.Sleep(1 * 1e9)
	}
}
