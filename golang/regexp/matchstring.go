package main

import (
	"fmt"
	"regexp"
)

func main() {
	matched, err := regexp.MatchString("foo.*", "seafood")
	fmt.Println(matched, err)
	matched, err = regexp.MatchString("bar.*", "seafood")
	fmt.Println(matched, err)
	matched, err = regexp.MatchString("a(b", "seafood")
	fmt.Println(matched, err)

	matched, err = regexp.MatchString("^[ _0-9a-zA-Z\u4E00-\u9FFF]+$", "~你好 ")
	fmt.Println(matched, err)
}
