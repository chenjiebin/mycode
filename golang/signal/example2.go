package main

import (
	"fmt"
	"os"
	"os/signal"
	//"syscall"
)

func main() {
	c := make(chan os.Signal)
	signal.Notify(c)
	//监听指定信号
	//signal.Notify(c, syscall.SIGHUP, syscall.SIGUSR2)

	//阻塞直至有信号传入
	s := <-c
	fmt.Println("get signal:", s)
}
