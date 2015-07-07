package action

var ActionFunc map[string]func(params map[string]string)

func init() {
	ActionFunc = make(map[string]func(params map[string]string))
}

func UserLogin() {

}
