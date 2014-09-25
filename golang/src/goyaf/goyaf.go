package goyaf

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"reflect"
	"strings"
	"time"
)

//goyaf instance
type Goyaf struct {
	version string
}

type GoyafController struct {
	ResponseWriter http.ResponseWriter
	Request        *Request
}

func (this *GoyafController) SetRequest(request *Request) {
	this.Request = request
}

func (this *GoyafController) GetRequest() *Request {
	return this.Request
}

type Request struct {
	Module     string
	Controller string
	Action     string
	r          *http.Request
}

func (this *Request) GetQuery(key string, defaultValue ...string) string {
	value := this.r.Form.Get(key)
	if len(value) == 0 && len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return value
}

func (this *Request) GetPost(key string) string {
	if this.r.Method != "POST" {
		return ""
	}

	this.r.ParseMultipartForm(32 << 20)
	if len(this.r.MultipartForm.Value[key]) > 0 {
		return this.r.MultipartForm.Value[key][0]
	}
	return ""
}

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

	Log(request)

	params := make([]reflect.Value, 1)
	params[0] = reflect.ValueOf(request)
	reflect.ValueOf(finalController).MethodByName("SetRequest").Call(params)

	reflect.ValueOf(finalController).MethodByName(strings.Title(uriSplits[3]) + "Action").Call(nil)

	//action, ok := Routes[r.URL.Path]
	//if ok {
	//	action()
	//	return
	//}

	//http.NotFound(w, r)
	//return
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
