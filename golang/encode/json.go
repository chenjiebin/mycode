package main

import (
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson"
	"reflect"
)

func main() {
	body := `{"errno":0,"errmsg":"10","data":{"member_id":"60006"}}`

	//使用simplejson进行解码
	js, err := simplejson.NewJson([]byte(body))

	if err != nil {
		panic(err.Error())
	}

	fmt.Println(js.Get("errno").Int())
	fmt.Println(js.Get("errmsg").String())
	//fmt.Println(js.Get("data")["member_id"])

	v := make(map[string]interface{})
	err = json.Unmarshal([]byte(body), &v)
	if err := json.Unmarshal([]byte(body), &v); err == nil {
		fmt.Println(v)
		//reflect.va
		i := reflect.ValueOf(v["data"]).Interface()
		a := i.(map[string]interface{})
		fmt.Println(a["member_id"])
	}

}
