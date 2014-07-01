package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	fmt.Println(8386725 % 20)
	fmt.Println("hello world")
	fmt.Println("The time is ", time.Now())

	fmt.Println("Or access the network:")
	fmt.Println(net.Dial("tcp", "pba.cn"))
}
