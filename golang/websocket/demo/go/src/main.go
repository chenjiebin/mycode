package main

import (
	"code.google.com/p/go.net/websocket"
	"container/list"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const (
	ErrMsgFormatError = 10001
	UserLoginSuccess  = 10010
	ErrUserIsExist    = 10011
)

var Tips = map[int]string{
	10001: "消息格式错误",
	10010: "用户登录成功",
	10011: "该用户已经存在了",
}

type User struct {
	Name   string
	Conn   *websocket.Conn
	ReadCh chan Message
	SendCh chan Message
	Quit   chan bool
}

//关闭与用户的连接
func (this *User) Close() {

}

//读取用户的输入
func (this *User) Read() {
	for {
		select {
		case message := <-this.ReadCh:
			fmt.Println(message)
		}
	}
}

//给用户发送消息
func (this *User) Send() {
	for {
		select {
		case message := <-this.SendCh:
			this.Conn.Write(JsonEncode(message))
			//case <-this.Quit:

			//	break
		}
	}
}

//用户集合存储
var Users *list.List

//检测用户是否存在
func IsExist(Name string) bool {
	for e := Users.Front(); e != nil; e = e.Next() {
		user := e.Value.(User)
		if user.Name == Name {
			return true
		}
	}
	return false
}

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

//tcp连接句柄，主要用来接受新用户的连接
func handle(ws *websocket.Conn) {
	var err error
	var user *User
	//设置用户名
	for {
		var reciveMsg string
		if err = websocket.Message.Receive(ws, &reciveMsg); err != nil {
			fmt.Println("Can't receive")
			return
		}
		message, err := JsonDecodeByString(reciveMsg)
		if err != nil {
			websocket.Message.Send(ws, "message error")
			continue
		}
		if IsExist(message.Content) {
			var returnMsg = Message{
				MessageType: ErrUserIsExist,
				FromUser:    "0",
				ToUser:      "",
				Content:     Tips[ErrUserIsExist],
				Time:        "",
			}
			ws.Write(JsonEncode(returnMsg))
			continue
		}
		user = &User{
			Name:   message.Content,
			Conn:   ws,
			ReadCh: make(chan Message),
			SendCh: make(chan Message),
			Quit:   make(chan bool),
		}
		Users.PushBack(*user)
		break
	}

	go user.Read()
	go user.Send()

	var returnMsg = Message{
		MessageType: UserLoginSuccess,
		FromUser:    "0",
		ToUser:      user.Name,
		Content:     Tips[UserLoginSuccess],
		Time:        "",
	}
	user.SendCh <- returnMsg

	for {
		var reciveMsg string
		if err = websocket.Message.Receive(ws, &reciveMsg); err != nil {
			fmt.Println("Can't receive")
			break
		}
		message, err := JsonDecodeByString(reciveMsg)
		if err != nil {
			var returnMsg = Message{
				MessageType: ErrMsgFormatError,
				FromUser:    "0",
				ToUser:      user.Name,
				Content:     Tips[ErrMsgFormatError],
				Time:        "",
			}
			user.SendCh <- returnMsg
			continue
		}
		user.ReadCh <- message
	}
}

func main() {
	Users = list.New()

	http.Handle("/", websocket.Handler(handle))
	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
