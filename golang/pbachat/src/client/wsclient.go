package client

import (
	"../common"
	"code.google.com/p/go.net/websocket"
	"encoding/json"
	"net"
)

type WSSocketClient struct {
	client
	Conn *websocket.Conn
}

var WSSocketClientList map[string]*WSSocketClient

func (wssc *WSSocketClient) Close() {
	wssc.Conn.Close()
	delete(SocketClientList, wssc.client.Id)
	common.Log("ws customer client num: ", len(WSSocketClientList))
}

func (wssc *WSSocketClient) Send(id string, msg []byte) {
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

func AddToList(socketClient *SocketClient) {
	if SocketClientList == nil {
		common.Log("init SocketClientList")
		SocketClientList = make(map[string]*SocketClient)
	}

	SocketClientList[socketClient.Id] = socketClient
	common.Log("customer total num: ", len(SocketClientList))
}
