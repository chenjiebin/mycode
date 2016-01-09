//一个channel有多个协程在监听，是哪个先处理
package main

import (
	"fmt"
	"time"
)

func foo(c chan int) {
	for {
		i := <-c
		fmt.Println("foo", i)
	}

}

func bar(c chan int) {
	for {
		i := <-c
		fmt.Println("bar", i)
	}
}

func main() {
	c := make(chan int)
	go foo(c)
	go bar(c)
	for i := 0; i < 100; i++ {
		c <- i
	}

	for {
		time.Sleep(1 * 1e9)
	}
}
