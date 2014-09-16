package goyaf

import (
	"log"
	"net/http"
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
	action, ok := Routes[r.URL.Path]
	if ok {
		action()
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
}
