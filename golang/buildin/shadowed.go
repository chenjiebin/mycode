// 模拟这种错误err is shadowed during return

package main

import (
	"strconv"
)

func main() {
	foo("100")
}

//func foo(i string) (err error) {
//	count := 0
//	if count == 0 {
//		_, err := strconv.ParseInt(i, 10, 64)
//		if err != nil {
//			return
//		}
//	}
//	return
//}

func foo(i string) (err error) {
	count := 0
	if count == 0 {
		_, err1 := strconv.ParseInt(i, 10, 64)
		if err1 != nil {
			err = err1
			return
		}
	}
	return
}
