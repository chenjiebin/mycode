package server

import (
	"../client"
	"../common"
	//"../csserver"
	"encoding/json"
	"net"
	"time"
)

type Message struct {
	Mtype   string
	Id      string
	Name    string
	Content string
}

func ClientHandler(conn net.Conn) {
	//实例化一个客户端，并增加进数组
	socketClient := &client.SocketClient{
		Conn: conn,
	}

	defer socketClient.Close()

	buffer := make([]byte, 2048)
	for {
		bytesRead, err := conn.Read(buffer)
		if err != nil {
			common.Log(err)
			break
		}

		msg := buffer[:bytesRead]
		common.Log("receive message from customer: ", msg)

		var jsonmsg Message
		json.Unmarshal(msg, &jsonmsg)
		common.Log("json decode: ", jsonmsg)

		switch jsonmsg.Mtype {
		case "10": //设置用户名
			socketClient.Id = jsonmsg.Id
			socketClient.Name = jsonmsg.Name
			client.AddToList(socketClient)
			socketClient.Send(jsonmsg.Id, []byte("set name success"))
		case "1": //发送给指定用户
			socketClient.Send(jsonmsg.Id, []byte(jsonmsg.Content))
		}
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
