package main

import (
	"fmt"
	"os"
	//"strconv"
)

func main() {
	write()
}

func write() {
	file, err := os.OpenFile("file2.txt", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		if os.IsNotExist(err) {
			os.Create("file2.txt")
			file, _ = os.OpenFile("file2.txt", os.O_APPEND|os.O_WRONLY, 0600)
		}
	}
	fmt.Println(os.IsNotExist(err))
	fmt.Println(file, err)

	file.Write([]byte("你好啊"))
	file.WriteString(",世界")
	//fmt.Println(file.WriteString(",世界"))

	////file.WriteAt([]byte("世界"), len([]byte("你好啊")))
	//fmt.Println(file)
}

func read() {
	file, err := os.Open("file.txt") // For read access.
	if err != nil {
		fmt.Println(err)
	}

	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("read %d bytes: %q\n", count, data[:count])
}
