package goyaf

//控制器集合
var Controllers map[string]interface{}

//增加控制器
func AddController(path string, controller interface{}) {
	Controllers[path] = controller
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

func init() {
	Controllers = make(map[string]interface{})
}
