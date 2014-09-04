//数据粘包问题

package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	netListen, err := net.Listen("tcp", ":9988")
	CheckError(err)

	defer netListen.Close()

	Log("Waiting for clients")
	for {
		conn, err := netListen.Accept()
		if err != nil {
			continue
		}

		Log(conn.RemoteAddr().String(), " tcp connect success")
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	// 消息缓冲
	msgbuf := bytes.NewBuffer(make([]byte, 0, 10240))
	// 数据缓冲
	databuf := make([]byte, 4096)
	// 消息长度
	length := 0
	// 消息长度uint32
	ulength := uint32(0)
	fmt.Print("ulength:", ulength)

	// 数据循环
	for {
		// 读取数据
		n, err := conn.Read(databuf)
		if err == io.EOF {
			fmt.Printf("Client exit: %s\n", conn.RemoteAddr())
		}
		if err != nil {
			fmt.Printf("Read error: %s\n", err)
			return
		}
		fmt.Println(databuf[:n])

		// 数据添加到消息缓冲
		n, err = msgbuf.Write(databuf[:n])
		if err != nil {
			fmt.Printf("Buffer write error: %s\n", err)
			return
		}

		// 消息分割循环
		for {
			// 消息头
			if length == 0 && msgbuf.Len() >= 4 {
				binary.Read(msgbuf, binary.LittleEndian, &ulength)
				length = int(ulength)
				// 检查超长消息
				if length > 10240 {
					fmt.Printf("Message too length: %d\n", length)
					return
				}
			}

			// 消息体
			if length > 0 && msgbuf.Len() >= length {
				fmt.Printf("Client messge: %s\n", string(msgbuf.Next(length)))
				length = 0
			} else {
				break
			}
		}
	}
}

func Log(v ...interface{}) {
	fmt.Println(v...)
}

func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
