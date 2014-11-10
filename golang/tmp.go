package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func a() (err error) {
	err = b()
	if err != nil {
		return
	}
	err = c()
	if err != nil {
		return
	}
	err = errors.New("a内错误")
	return
}

func b() (err error) {
	err = errors.New("b内错误")
	return
}

func c() (err error) {
	err = errors.New("c内错误")
	return
}

func main() {
	err := a()
	if err != nil {
		fmt.Println(err)
	}
}
