package goyaf

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strings"
	"time"
)

type Goyaf struct {
	version string
}

//控制器对象
type GoyafController struct {
	Response *Response
	Request  *Request
}

func (this *GoyafController) SetRequest(request *Request) {
	this.Request = request
}

func (this *GoyafController) GetRequest() *Request {
	return this.Request
}

func (this *GoyafController) SetResponse(response *Response) {
	this.Response = response
}

func (this *GoyafController) GetResponse() *Response {
	return this.Response
}

//请求对象
type Request struct {
	Module     string
	Controller string
	Action     string
	r          *http.Request
}

func (this *Request) GetQuery(key string, defaultValue ...string) string {
	this.r.ParseForm()
	value := this.r.Form.Get(key)
	if len(value) == 0 && len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return value
}

//获取所有的参数
//func (this *Request) GetQuerys() string {
//	this.r.ParseForm()
//	value := this.r.Form.Get(key)
//	if len(value) == 0 && len(defaultValue) > 0 {
//		return defaultValue[0]
//	}
//	return value
//}

func (this *Request) GetPost(key string, defaultValue ...string) string {
	if this.r.Method != "POST" {
		return ""
	}

	this.r.ParseMultipartForm(32 << 20)
	values := this.r.MultipartForm.Value[key]
	if len(values) > 0 && len(values[0]) > 0 {
		return values[0]
	}

	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return ""
}

func (this *Request) GetPosts() url.Values {
	//todo 关于Content-Type: multipart/form-data;还需处理
	this.r.ParseForm()
	return this.r.PostForm
}

func (this *Request) GetCookie(key string, defaultValue ...string) string {
	//todo 这里还要测试cookie的有效性
	cookie, err := this.r.Cookie(key)
	if err != nil {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return ""
	}

	return cookie.Value
}

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

//默认路由
type GoyafMux struct{}

func (p *GoyafMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	uriSplits := strings.Split(r.RequestURI, "/")
	if len(uriSplits) < 4 {
		http.NotFound(w, r)
		return
	}

	is404 := true
	var finalController interface{}
	for path, controller := range Controllers {
		if strings.Index(r.RequestURI, path) == 0 {
			finalController = controller
			is404 = false
			break
		}
	}

	if is404 {
		http.NotFound(w, r)
		return
	}

	request := &Request{
		Module:     uriSplits[1],
		Controller: uriSplits[2],
		Action:     uriSplits[3],
		r:          r,
	}
	params := make([]reflect.Value, 1)
	params[0] = reflect.ValueOf(request)
	reflect.ValueOf(finalController).MethodByName("SetRequest").Call(params)

	response := &Response{
		w: w,
	}
	responseParams := make([]reflect.Value, 1)
	responseParams[0] = reflect.ValueOf(response)
	reflect.ValueOf(finalController).MethodByName("SetResponse").Call(responseParams)

	Debug(strings.Title(uriSplits[3]) + "Action")
	action := reflect.ValueOf(finalController).MethodByName(strings.Title(uriSplits[3]) + "Action")
	Debug(action)
	if action.IsValid() {
		action.Call(nil)
		response.Response()
		return
	}

	http.NotFound(w, r)
	return
}

var Routes map[string]func()
var Controllers map[string]interface{}

func AddController(path string, controller interface{}) {
	Controllers[path] = controller
}

func AddRoute(path string, action func()) {
	Routes[path] = action
}

func init() {
	Routes = make(map[string]func())
	Controllers = make(map[string]interface{})
}

func Run() {
	mux := &GoyafMux{}

	err := http.ListenAndServe(":9090", mux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	Log("start server success.")
}

//日志记录
func Log(v ...interface{}) {
	systemTime := time.Now().Format("2006-01-02 15:04:05")

	fmt.Print(systemTime, " Log: ")
	fmt.Println(v...)
}

//调试信息
func Debug(v ...interface{}) {
	systemTime := time.Now().Format("2006-01-02 15:04:05")

	fmt.Print(systemTime, " Debug: ")
	fmt.Println(v...)
}

//检查错误
func CheckError(err error) {
	if err != nil {
		systemTime := time.Now().Format("2006-01-02 15:04:05")

		fmt.Fprintf(os.Stderr, "%s Fatal error: %s", systemTime, err.Error())
		os.Exit(1)
	}
}
