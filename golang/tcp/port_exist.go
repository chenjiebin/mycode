package main

import (
	"fmt"
	"net"
)

func main() {

	tcpAddr, err := net.ResolveTCPAddr("tcp4", ":8000")
	fmt.Println(tcpAddr, err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	net.DialIP()
	fmt.Println(conn, err)

	defer conn.Close()

}
