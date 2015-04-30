package main

import (
	"fmt"
	"sync"
	"time"
)

//在平时使用锁中，非常容易忘记关闭锁，GO语言处理这就非常简单，使用defer就可以很好避免这个问题
func write() {
	var mutex sync.Mutex
	mutex.Lock()
	defer mutex.Unlock()
}

func repeatedlyLock() {
	var mutex sync.Mutex
	fmt.Println("Lock the lock. (G0)")
	//线程G0锁定
	mutex.Lock()
	fmt.Println("The lock is locked. (G0)")

	//启动三个协程，每个协程在调用mutex.Lock时，会因为G0线程处于锁状态，而处于阻塞状态
	for i := 1; i <= 3; i++ {
		go func(i int) {
			fmt.Printf("Lock the lock. (G%d)\n", i)
			mutex.Lock()
			fmt.Printf("The lock is locked. (G%d)\n", i)
		}(i)
	}
	time.Sleep(time.Second)
	fmt.Println("Unlock the lock. (G0)")
	//线程G0解锁后
	mutex.Unlock()
	fmt.Println("The lock is unlocked. (G0)")
	time.Sleep(time.Second)
}

func main() {
	repeatedlyLock()
}
