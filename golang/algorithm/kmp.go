// 用代码实现kmp

package main

import (
	"fmt"
)

func main() {
	kmp()
}

func normal() {
	l := "BBC ABCDAB ABCDABCDABDE"
	s := "ABCDABD"

	lLen := len(l)
	sLen := len(s)

	for i := 0; i < lLen; i++ {
		var j int
		for j = 0; j < sLen; j++ {
			if s[j] != l[i+j] {
				break
			}
		}
		// 说明已经在长串中找到短串
		if j == sLen {
			fmt.Println(i)
		}
	}

}

func kmp() {
	l := "BBC ABCDAB ABCDABCDABDE"
	s := "ABCDABD"

	lLen := len(l)
	sLen := len(s)

	// 计算s字符串的每位匹配值
	next := make([]int, sLen) // 声明一个数组来存放字符串s的每位匹配值
	next[0] = 0
	k := 0 // 匹配值数值
	for i := 1; i < sLen; i++ {
		if s[i] == s[k] {
			k++
		} else {
			k = 0
		}
		next[i] = k
	}
	fmt.Println(s)
	fmt.Println(next)

	var j int
	for i := 0; i < lLen; {
		for j = 0; j < sLen; j++ {
			if s[j] != l[i+j] {
				break
			}
		}
		// 说明已经在长串中找到短串
		if j == sLen {
			fmt.Println(i)
		}

		// 如果首个字符就不符合，则长串直接到下一位进行比较
		if j == 0 {
			i++
			continue
		}

		i = i + j - next[j-1]

	}
}
