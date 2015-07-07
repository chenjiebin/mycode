//聊天服务端
//剩余问题
//1.UserList并发问题
//2.单条发送信息问题
//3.MessageType问题
package main

import (
	"./message"
	"code.google.com/p/go.net/websocket"
	"log"
	"net/http"
	"os"
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
	ReadCh chan message.Message
	SendCh chan message.Message
}

//关闭与用户的连接
func (this *User) Close() {
	if len(this.Name) == 0 {
		this.Conn.Close()
		return
	}
	UserList.RemoveUser(this)
	this.Conn.Close()
	Debug(this, "关闭了连接")
}

//读取用户的输入
func (this *User) Read() {
	for {
		select {
		case msg := <-this.ReadCh:
			Debug("收到客户端", this, "发送的信息:", msg)

			switch msg.MessageType {
			case 1: //设置用户名
				if len(this.Name) > 0 {
					continue
				}
				if UserList.IsExist(msg.Content) {
					this.SendCh <- message.Message{
						MessageType: ErrUserIsExist,
						FromUser:    "0",
						ToUser:      "",
						Content:     Tips[ErrUserIsExist],
						Time:        "",
					}
					continue
				}
				this.Name = msg.Content
				UserList.Add(this)

				this.SendCh <- message.Message{
					MessageType: UserLoginSuccess,
					FromUser:    "0",
					ToUser:      this.Name,
					Content:     Tips[UserLoginSuccess],
					Time:        "",
				}
				Debug(this, "设置用户名成功")
			case 2: //群发消息
				for _, u := range UserList.Users {
					u.SendCh <- message.Message{
						MessageType: 2,
						FromUser:    this.Name,
						ToUser:      u.Name,
						Content:     msg.Content,
						Time:        "",
					}
				}
			}
		}
	}
}

//给用户发送消息
func (this *User) Send() {
	for {
		select {
		case msg := <-this.SendCh:
			Debug("给客户端", this, "发送消息:", msg)
			this.Conn.Write(message.JsonEncode(msg))
		}
	}
}

//用户群对象
var UserList Users

//用户群
type Users struct {
	Users map[string]*User
}

//增加用户
func (this *Users) Add(u *User) {
	if this.IsExist(u.Name) {
		return
	}
	this.Users[u.Name] = u
}

//检测用户是否存在
func (this *Users) IsExist(Name string) bool {
	_, ok := this.Users[Name]
	return ok
}

//移除用户
func (this *Users) RemoveUser(u *User) {
	delete(this.Users, u.Name)
}

//tcp连接句柄，主要用来接受新用户的连接
func handle(ws *websocket.Conn) {
	var err error
	var user *User
	user = &User{
		Name:   "",
		Conn:   ws,
		ReadCh: make(chan message.Message),
		SendCh: make(chan message.Message),
	}
	Debug("用户连上线了:", user)

	go user.Read()
	go user.Send()

	for {
		var reciveMsg string
		if err = websocket.Message.Receive(ws, &reciveMsg); err != nil {
			user.Close()
			break
		}
		msg, err := message.JsonDecodeByString(reciveMsg)
		if err != nil {
			var returnMsg = message.Message{
				MessageType: ErrMsgFormatError,
				FromUser:    "0",
				ToUser:      user.Name,
				Content:     Tips[ErrMsgFormatError],
				Time:        "",
			}
			user.SendCh <- returnMsg
			continue
		}
		user.ReadCh <- msg
	}
}

var DebugLog *log.Logger

func Debug(v ...interface{}) {
	DebugLog.Println(v...)
}

func main() {
	UserList = Users{
		Users: make(map[string]*User),
	}

	DebugLog = log.New(os.Stdout, "[DEBUG]", log.LstdFlags)

	http.Handle("/", websocket.Handler(handle))
	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
