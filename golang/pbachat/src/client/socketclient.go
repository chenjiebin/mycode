package client

import (
	"../common"
	"encoding/json"
	"net"
)

type SocketClient struct {
	client
	Conn net.Conn
}

var SocketClientList map[string]*SocketClient

func (sc *SocketClient) Close() {
	sc.Conn.Close()
	delete(SocketClientList, sc.client.Id)
	common.Log("customer client num: ", len(SocketClientList))
}

func (sc *SocketClient) Send(id string, msg []byte) {
	socketClient, ok := SocketClientList[id]
	if ok {
		tmpmsg := message{"1", sc.Id, sc.Name, socketClient.Id, socketClient.Name, msg}
		common.Log(tmpmsg)

		jsonmsg, _ := json.Marshal(tmpmsg)
		socketClient.Conn.Write(jsonmsg)
	} else {
		sc.Conn.Write([]byte("user is not exist. user_id: " + id))
	}
}

func (sc *SocketClient) SendToAll(msg []byte) {
	for _, c := range SocketClientList {
		common.Log("send message to customer: ", c.Conn)
		c.Conn.Write(msg)
	}
}

func AddToList(socketClient *SocketClient) {
	if SocketClientList == nil {
		common.Log("init SocketClientList")
		SocketClientList = make(map[string]*SocketClient)
	}

	SocketClientList[socketClient.Id] = socketClient
	common.Log("customer total num: ", len(SocketClientList))
}
