package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGUSR2)

	//当调用了该方法后，下面的for循环内<-c接收到一个信号就退出了。
	signal.Stop(c)

	for {
		s := <-c
		fmt.Println("get signal:", s)
	}

}
