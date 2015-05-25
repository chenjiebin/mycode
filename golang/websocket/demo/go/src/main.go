//websocket服务器
package main

import (
	"code.google.com/p/go.net/websocket"
	"container/list"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Name string
	Conn *websocket.Conn
	In   chan string
	Out  chan string
	Quit bool
}

//关闭与用户的连接
func (this *User) Close() {

}

//读取用户的输入
func (this *User) Read() {

}

//给用户发送消息
func (this *User) Send() {
	for {
		select {
		case buffer := <-this.Incoming:
			this.Conn.Write([]byte(buffer))
		case <-this.Quit:
			client.Conn.Close()
			break
		}
	}
}

//用户组
var users *list.List

//消息结构
type Message struct {
	MessageType int
	FromUser    string
	ToUser      string
	Content     string
	Time        string
}

//消息json编码成字节
func JsonEncode(message Message) []byte {
	byteMessage, _ := json.Marshal(message)
	return byteMessage
}

//二进制消息json解码成结构体
func JsonDecode(byteMessage []byte) (message Message, err error) {
	err = json.Unmarshal(byteMessage, &message)
	if err != nil {
		return
	}
	return
}

//字符串消息json解码成结构体
func JsonDecodeByString(strmsg string) (message Message, err error) {
	message, err = JsonDecode([]byte(strmsg))
	if err != nil {
		return
	}
	return
}

func handle(ws *websocket.Conn) {
	var err error
	user := User{
		Conn: ws,
	}
	users.PushBack(user)

	go user.Read()
	go user.Send()

	for {
		var reciveMsg string
		if err = websocket.Message.Receive(ws, &reciveMsg); err != nil {
			fmt.Println("Can't receive")
			break
		}
		user.Out <- 
		
		message, err := JsonDecodeByString(reciveMsg)
		if err != nil {
			websocket.Message.Send(ws, "message error")
		}

		switch message {
		case 1:
			user.Name = message.Content

		case 2:
		}

		//		for _, c := range users {
		//			if err = websocket.Message.Send(c.Conn, msg); err != nil {
		//				fmt.Println("Can't send")
		//				break
		//			}
		//		}
	}
}

func main() {
	http.Handle("/", websocket.Handler(handle))
	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
