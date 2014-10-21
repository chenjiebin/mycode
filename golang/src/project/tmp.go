package main

import (
	"fmt"
)

func main() {
	params := map[string]string{
		"page": "",
		"name": "cjb",
	}
	checkPageAndCount(&params)

	fmt.Println(params)
}

//检测page和count参数，如果没有值赋予默认值1和10
//因为page和count参数都是在列表的时候才需要传递，所以这里直接计算出offset
func checkPageAndCount(params *map[string]string) {
	(*params)["page"] = "1"

}
