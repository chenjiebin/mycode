package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"syscall"
	"time"
)

var c chan int

func main() {
	c = make(chan int, 10)
	c <- 10
	c <- 9

	ret, ret2, err := syscall.RawSyscall(syscall.SYS_FORK, 0, 0, 0)

	if ret2 == 0 {
		fmt.Println(ret, ret2, err)
		fmt.Println("child process")
		go listenChan(c)
	} else {
		fmt.Println(ret, ret2, err)
		fmt.Println("parent process")
		go startByServer(c)
	}

	for {
		time.Sleep(1 * 1e9)
	}
}

func listenChan(c chan int) {
	fmt.Println(<-c)
	for {
		select {
		case i := <-c:
			fmt.Println(i)
		}
	}

}

func startByServer(c chan int) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		c <- 8
		fmt.Println(<-c)
		fmt.Fprintf(w, "Hello world")
	})

	l, err := net.Listen("tcp", ":9090")
	if nil != err {
		log.Fatalln(err)
	}

	err = http.Serve(l, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
