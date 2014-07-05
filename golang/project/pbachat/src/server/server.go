package server

import (
	"../client"
	"../common"
	//"../csserver"
	"net"
	"time"
)

func ClientHandler(conn net.Conn) {
	//实例化一个客户端，并增加进数组
	newClient := &client.Client{
		ConnType: 1,
		Conn:     conn,
	}
	client.ClientList.PushBack(*newClient)
	common.Log(client.ClientList.Len())

	defer newClient.Close()

	buffer := make([]byte, 2048)
	for {
		bytesRead, err := conn.Read(buffer)
		if err != nil {
			common.Log(err)
			newClient.Close()
			break
		}

		common.Log(string(buffer[:bytesRead]))

		newClient.SendToAll(buffer[:bytesRead])
	}
}

func StartServer() {
	netListen, err := net.Listen("tcp", ":9988")
	common.CheckError(err)

	//defer函数退出时执行
	defer netListen.Close()

	for {
		common.Log("Waiting for clients")
		conn, err := netListen.Accept()
		if err != nil {
			continue
		}
		//连线上的是否，返回当前系统时间
		conn.Write([]byte(time.Now().String()))

		go ClientHandler(conn)
	}
}
