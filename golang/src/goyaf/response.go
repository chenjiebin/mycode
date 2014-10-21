package goyaf

import (
	"fmt"
	"net/http"
)

//返回对象
type Response struct {
	w     http.ResponseWriter
	bodys []interface{}
}

//追加body
func (this *Response) AppendBody(body interface{}) {
	this.bodys = append(this.bodys, body)
}

//返回数据
func (this *Response) Response() {
	for _, body := range this.bodys {
		fmt.Fprintln(this.w, body)
	}
}

//设置输出头
func (this *Response) SetHeader(key string, value string) {
	this.w.Header().Set(key, value)
}

//设置cookie
func (this *Response) SetCookie(cookie *http.Cookie) {
	http.SetCookie(this.w, cookie)
}
