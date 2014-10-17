package goyaf

import (
	"fmt"
	"net/http"
)

//返回对象
type Response struct {
	w    http.ResponseWriter
	body interface{}
}

func (this *Response) SetBody(body interface{}) {
	this.body = body
}

func (this *Response) GetBody() interface{} {
	return this.body
}

func (this *Response) Response() {
	fmt.Fprintln(this.w, this.body)
}

func (this *Response) SetCookie(cookie *http.Cookie) {
	http.SetCookie(this.w, cookie)
}
