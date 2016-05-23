// golang实现循环移位
// 例如abcdefg，将该字符串循环向左移动两位就变成defgab
// 要求实现的复杂度是O(n)，空间复杂度是O(1)
package main

import (
	"fmt"
)

func main() {
	s := "abcdefg"
	fmt.Println(reverseString(s, 0, 6))

}

//// 常用的操作方法
//func common() {
//	s := "abcdefg"
//	fmt.Println(reverseString(s))
//}

// 反转字符串
func reverseString(s string, from int, to int) string {
	runes := []rune(s)
	for ; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}
