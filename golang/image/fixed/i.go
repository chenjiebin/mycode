//验证fixed.i函数int26_6转换成int
package main

import (
	"fmt"
	"golang.org/x/image/math/fixed"
)

func main() {
	fmt.Println(fixed.I(3))
	fmt.Println(int(fixed.I(3)) >> 6)
}
