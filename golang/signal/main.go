package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	//设置一个channel来发送信号
	c := make(chan os.Signal, 1)
	signal.Notify(c)

	// 一直运行一直到收到一个信号

	s := <-c

	fmt.Println("Got signal:", s) //当我停止运行时 Got signal: interrupt

}
