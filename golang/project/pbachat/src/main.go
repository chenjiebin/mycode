package main

import (
	"container/list"
	"fmt"
	"net"
	"os"
	"time"
)

var clientList list.List

type Client struct {
	name string
	Conn net.Conn
}

//关闭用户连接，并且从用户数组中移除该用户
func (client *Client) Close() {
	client.Conn.Close()

	for entry := clientList.Front(); entry != nil; entry = entry.Next() {
		if client.Conn == entry.Value.(Client).Conn {
			Log("remove from clientList", client.Conn)
			clientList.Remove(entry)
		}
	}
	Log(clientList.Len())
}

func ClientHandler(conn net.Conn) {
	//实例化一个客户端，并增加进数组
	newClient := &Client{
		Conn: conn,
	}
	clientList.PushBack(*newClient)
	Log(clientList.Len())

	buffer := make([]byte, 2048)
	for {
		bytesRead, err := conn.Read(buffer)
		if err != nil {
			Log(err)
			newClient.Close()
			break
		}

		Log(string(buffer[:bytesRead]))

		//群发消息
		for entry := clientList.Front(); entry != nil; entry = entry.Next() {
			Log(entry.Value.(Client).Conn)
			entry.Value.(Client).Conn.Write(buffer[:bytesRead])
		}
	}
}

func startServer() {
	netListen, err := net.Listen("tcp", ":9988")
	checkError(err)

	//defer函数退出时执行
	defer netListen.Close()

	for {
		Log("Waiting for clients")
		conn, err := netListen.Accept()
		if err != nil {
			continue
		}
		//连线上的是否，返回当前系统时间
		conn.Write([]byte(time.Now().String()))

		go ClientHandler(conn)
	}
}

func main() {
	startServer()
}

func Log(v ...interface{}) {
	fmt.Println(v...)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
