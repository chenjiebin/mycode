package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strconv"
)

func main() {
	var i int64
	i = 0x100
	fmt.Println(int64ToString(i))

	//fmt.Println(dlIntToByte(125))
}

func int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
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
