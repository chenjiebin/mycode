package error

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
