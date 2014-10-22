package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

var running bool

var conn net.Conn

func read(conn net.Conn) {
	for {
		buffer := make([]byte, 2048)
		bytesRead, err := conn.Read(buffer)
		checkError(err)

		fmt.Println("receive: ", string(buffer[:bytesRead]))
	}
}

func write(conn net.Conn) {
	for {
		fmt.Print("> ")
		reader := bufio.NewReader(os.Stdin)
		words, _ := reader.ReadBytes('\n')

		conn.Write([]byte(words))
	}
}

func main() {
	running = true

	service := "127.0.0.1:9988"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	defer conn.Close()

	Log("connect success")

	go read(conn)
	go write(conn)

	for running {
		time.Sleep(1 * 1e9)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func Log(v ...interface{}) {
	fmt.Println(v...)
}
