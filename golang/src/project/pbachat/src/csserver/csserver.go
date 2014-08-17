package csserver

import (
	"../client"
	"../common"
	"code.google.com/p/go.net/websocket"
	"fmt"
	"log"
	"net/http"
)

func clientHandle(wsconn *websocket.Conn) {
	var err error

	//实例化一个客户端，并增加进数组
	newClient := &client.Client{
		ConnType: 2,
		WsConn:   wsconn,
	}
	client.ClientList.PushBack(*newClient)
	common.Log(client.ClientList.Len())

	defer newClient.Close()

	for {
		var reply string

		if err = websocket.Message.Receive(wsconn, &reply); err != nil {
			newClient.Close()
			break
		}

		fmt.Println("Received back from client: " + reply)

		msg := "Received: " + reply
		fmt.Println("Sending to client: " + msg)

		newClient.SendToAll([]byte(msg))
	}
}

func StartCSClientServer() {
	http.Handle("/", websocket.Handler(clientHandle))
	if err := http.ListenAndServe(":9999", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
