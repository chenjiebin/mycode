//聊天服务端
//剩余问题
//1.UserList并发问题
//2.单条发送信息问题
//3.MessageType问题
package main

import (
	ourErr "./error"
	"./log"
	"./message"
	"./user"
	"code.google.com/p/go.net/websocket"
	"net/http"
)

//tcp连接句柄，主要用来接受新用户的连接
func handle(ws *websocket.Conn) {
	var err error
	var u *user.User
	u = &user.User{
		Name:   "",
		Conn:   ws,
		ReadCh: make(chan message.Message),
		SendCh: make(chan message.Message),
	}
	log.Debug("用户连上线了:", u)

	go u.Read()
	go u.Send()

	for {
		var reciveMsg string
		if err = websocket.Message.Receive(ws, &reciveMsg); err != nil {
			u.Close()
			break
		}

		log.Debug("收到客户端发送的原始消息:", reciveMsg)

		msg, err := message.JsonDecodeByString(reciveMsg)
		if err != nil {
			var returnMsg = message.Message{
				MessageType: ourErr.ErrMsgFormatError,
				FromUser:    "0",
				ToUser:      u.Name,
				Content:     ourErr.Tips[ourErr.ErrMsgFormatError],
				Time:        "",
			}
			u.SendCh <- returnMsg
			continue
		}
		u.ReadCh <- msg
	}
}

func main() {

	http.Handle("/", websocket.Handler(handle))
	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
