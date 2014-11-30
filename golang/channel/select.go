package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 10)

	go hello(c)

	for {

		i := <-c
		fmt.Println(i)
		if i == 3 {
			break
		}

		//select {
		//case i := <-c:
		//	fmt.Println(i)
		//	if i == 3 {
		//		break
		//	}
		//default:
		//	break
		//}
	}

	//time.Sleep(10 * 1e9)
}

func hello(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
		time.Sleep(1 * 1e9)
	}

}
