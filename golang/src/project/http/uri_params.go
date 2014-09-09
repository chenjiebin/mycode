package main

import (
	"fmt"
	"strings"
)

//分析如/page/1/count/10/这样的参数
func analyseStringParams(stringParams string) map[string]string {
	params := make(map[string]string)

	tmp := strings.Split(strings.Trim(stringParams, "/"), "/")
	maxRange := len(tmp) - 1
	for i := 0; i < maxRange; i = i + 2 {
		params[tmp[i]] = tmp[i+1]
	}

	return params
}

func main() {
	fmt.Println(analyseStringParams("/page/1/count/10/name"))
}
