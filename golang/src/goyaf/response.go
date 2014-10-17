package goyaf

import (
	"fmt"
	"net/http"
)

//返回对象
type Response struct {
	w     http.ResponseWriter
	body  interface{}
	bodys []interface{}
}

func (this *Response) SetBody(body interface{}) {
	this.body = body
}

func (this *Response) GetBody() interface{} {
	return this.body
}

func (this *Response) AppendBody(body interface{}) {
	this.bodys = append(this.bodys, body)
}

func (this *Response) Response() {
	fmt.Fprintln(this.w, this.bodys)
}

func (this *Response) SetCookie(cookie *http.Cookie) {
	http.SetCookie(this.w, cookie)
}
