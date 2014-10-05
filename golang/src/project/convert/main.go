package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strconv"
)

func main() {
	i, _ := strconv.Atoi("100")
	fmt.Println(i)
}

func int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

func intToString() {
	fmt.Println(strconv.Itoa(100))
}

func stringToInt() {
	i, _ := strconv.Atoi("100")
	fmt.Println(i)
}

func dlIntToByte(i int) []byte {
	iStr := strconv.Itoa(i)
	zeroLength := 9 - len(iStr)
	for i := 0; i < zeroLength; i = i + 1 {
		iStr = "0" + iStr
	}
	return []byte(iStr)
}

func bytesToInt32() {
	b := []byte{0x00, 0x00, 0x03, 0xe8}
	bytesBuffer := bytes.NewBuffer(b)

	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)
	fmt.Println(x)
}

func int32ToBytes() {
	var x int32
	x = 106
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	fmt.Println(bytesBuffer.Bytes())
}

func byteToString() {
	fmt.Println(string([]byte{97, 98, 99, 100}))
}

func strintToBytes() {
	fmt.Println([]byte("abcd"))
}
