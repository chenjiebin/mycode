package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"syscall"
	"time"
)

func main() {
	var ret, ret2 uintptr
	var err syscall.Errno

	darwin := runtime.GOOS == "darwin"

	// fork off the parent process
	ret, ret2, err = syscall.RawSyscall(syscall.SYS_FORK, 0, 0, 0)
	if err != 0 {
		log.Fatalln(err)
	}

	// failure
	if ret2 < 0 {
		os.Exit(-1)
	}

	// handle exception for darwin
	if darwin && ret2 == 1 {
		ret = 0
	}

	if ret > 0 {
		fmt.Println("parent process")
	} else {
		fmt.Println("child process")
		fmt.Println(syscall.Getpid())
		fmt.Println(syscall.Getppid())
	}

	for {
		time.Sleep(1 * 1e9)
	}
}
