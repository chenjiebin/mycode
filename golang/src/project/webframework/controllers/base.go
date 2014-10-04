//控制器基类
package controllers

import (
	"encoding/json"
	"goyaf"
)

type Base struct {
	goyaf.GoyafController
}

//输出成功信息
func (this *Base) printSuccessMessage(message, data interface{}) {
	result := make(map[string]interface{})
	result["errno"] = "0"
	result["errmsg"] = ""
	result["data"] = data

	jsonResult, _ := json.Marshal(result)

	isJsonp := this.GetRequest().GetQuery("is_jsonp")
	if len(isJsonp) == 0 {
		this.GetResponse().SetBody(string(jsonResult))
	} else {
		callback := this.GetRequest().GetQuery("callback")
		this.GetResponse().SetBody(callback + "(" + string(jsonResult) + ")")
	}
}
