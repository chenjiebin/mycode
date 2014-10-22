//websocket服务器
package main

import (
	"code.google.com/p/go.net/websocket"
	"fmt"
	"log"
	"net/http"
)

type Client struct {
	Conn *websocket.Conn
}

var clientList []Client

func Echo(ws *websocket.Conn) {
	var err error

	client := Client{
		Conn: ws,
	}

	clientList = append(clientList, client)

	for {
		var reply string

		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't receive")
			break
		}

		fmt.Println("Received back from client: " + reply)

		msg := "Received: " + reply
		fmt.Println("Sending to client: " + msg)

		for _, c := range clientList {
			if err = websocket.Message.Send(c.Conn, msg); err != nil {
				fmt.Println("Can't send")
				break
			}
		}
	}
}

func main() {
	http.Handle("/", websocket.Handler(Echo))
	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
