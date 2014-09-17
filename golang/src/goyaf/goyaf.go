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
	Request        *http.Request
}

type GoyafMux struct{}

func (p *GoyafMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.RequestURI)

	uriSplits := strings.Split(r.RequestURI, "/")
	if len(uriSplits) < 4 {
		http.NotFound(w, r)
		return
	}

	is404 := true
	var finalController interface{}
	for path, controller := range Controllers {
		log.Println(path, controller)
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

	action = strings.Split(r.RequestURI, "/")

	reflect.ValueOf(finalController).MethodByName("TestAction").Call(nil)

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
