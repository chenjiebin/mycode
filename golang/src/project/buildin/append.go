package main

import (
	"fmt"
)

func main() {

}

func appendBytes() {
	str := "hehe"

	strBytes := []byte(str)
	intBytes := []byte{0x00, 0x00, 0x03, 0xe8}

	fmt.Println(append(strBytes[0:], intBytes[0:]...))
}
