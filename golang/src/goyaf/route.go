package goyaf

import (
	"net/http"
	"reflect"
	"strings"
)

var Routes map[string]func()

func init() {
	Routes = make(map[string]func())
}

//默认路由
type GoyafMux struct{}

func (p *GoyafMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//记录请求
	Log("access: " + r.RemoteAddr + " " + r.Method + " " + r.RequestURI)

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
		if panicHandle != nil {
			defer panicHandle()
		}

		action.Call(nil)
		response.Response()
		return
	}

	http.NotFound(w, r)
	return
}
