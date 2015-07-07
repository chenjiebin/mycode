package main

import (
	"./message"
	"fmt"
)

func main() {
	msg, _ := message.JsonDecodeByString(`{"MessageType":1,"FromUser":"","ToUser":"","Content":"{\"MessageType\":1,\"FromUser\":\"\",\"ToUser\":\"\",\"Content\":\"\\u9648\\u6770\\u658c\",\"Time\":\"\"}","Time":""}`)
	fmt.Println(msg)
	msg2, _ := message.JsonDecodeByString(msg.Content)
	fmt.Println(msg2)
}
