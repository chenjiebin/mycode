package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	write()
	read()
}

func write() {
	err := ioutil.WriteFile("./ioutil.txt", []byte("Hello world!"), 0777)
	fmt.Println(err)
}

func read() {
	data, err := ioutil.ReadFile("./ioutil2.txt")
	if os.IsNotExist(err) {
		fmt.Println("file not exist")
		return
	}
	fmt.Println(string(data), err)
}
