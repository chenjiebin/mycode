package user

import (
	"../error"
	"../log"
	"../message"
	"code.google.com/p/go.net/websocket"
)

//用户结构体
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
	log.Debug(this, "关闭了连接")
}

//读取用户的输入
func (this *User) Read() {
	for {
		select {
		case msg := <-this.ReadCh:
			log.Debug("收到客户端", this, "发送的信息:", msg)

			switch msg.MessageType {
			case 1: //设置用户名
				if len(this.Name) > 0 {
					continue
				}
				if UserList.IsExist(msg.Content) {
					this.SendCh <- message.Message{
						MessageType: error.ErrUserIsExist,
						FromUser:    "0",
						ToUser:      "",
						Content:     error.Tips[error.ErrUserIsExist],
						Time:        "",
					}
					continue
				}
				this.Name = msg.Content
				UserList.Add(this)

				this.SendCh <- message.Message{
					MessageType: error.UserLoginSuccess,
					FromUser:    "0",
					ToUser:      this.Name,
					Content:     error.Tips[error.UserLoginSuccess],
					Time:        "",
				}
				log.Debug(this, "设置用户名成功")
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
			log.Debug("给客户端", this, "发送消息:", msg)
			this.Conn.Write(message.JsonEncode(msg))
		}
	}
}
