package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	h := md5.New()
	h.Write([]byte("123456"))
	fmt.Printf("%s\n", hex.EncodeToString(h.Sum(nil)))
}
