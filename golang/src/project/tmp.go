package main

import (
	"fmt"
)

func main() {
	name := "哈ab哈哈哈"
	ru := []rune(name)

	fmt.Println(len(ru))

	for _, r := range ru {
		fmt.Println(string(r))
	}
	//fmt.Println(len("爸爸爸爸爸爸爸爸爸爸爸爸爸爸爸爸爸爸爸爸爸爸爸爸爸爸爸爸爸爸爸爸爸爸爸爸爸爸爸爸爸爸爸爸爸"))
	//params := map[string]string{
	//	"page": "",
	//	"name": "cjb",
	//}
	//checkPageAndCount(&params)
	//爸爸爸爸爸爸爸爸爸爸爸爸爸爸爸
	//fmt.Println(params)
	//爸爸爸爸爸爸爸爸爸爸爸爸爸爸爸爸爸爸爸爸爸爸爸爸爸爸爸爸爸爸爸爸爸爸爸爸爸爸爸爸爸爸爸爸爸爸爸爸爸爸爸爸爸爸爸
}

//检测page和count参数，如果没有值赋予默认值1和10
//因为page和count参数都是在列表的时候才需要传递，所以这里直接计算出offset
func checkPageAndCount(params *map[string]string) {
	(*params)["page"] = "1"

}
