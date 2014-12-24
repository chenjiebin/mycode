package main

import (
	"fmt"
	//"runtime"
	"time"
)

func say(s string) {
	for i := 0; i < 1; i++ {
		//runtime.Gosched()
		fmt.Println(s)
	}
}

func calc() {
	j := 0
	for i := 0; i < 100; i++ {
		j += 1
	}
	time.Sleep(1 * 1e9)
	fmt.Println(j)
}

func main() {
	gogo()

	for {
		time.Sleep(1 * 1e9)
	}
	//go say("world")
	//say("hello")
}

func gogo() {
	say("Hello")
	go calc()
}
