package action

var ActionFunc map[string]func(params map[string]string)

func init() {
	ActionFunc = make(map[string]func(params map[string]string))
}

//用户登录功能
func UserLogin(params map[string]string) {
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
}
