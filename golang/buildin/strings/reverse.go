package main

// 反转字符串
func reverseString(s string, from int, to int) string {
	runes := []rune(s)
	for ; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}
