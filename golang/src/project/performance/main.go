//+=性能测试
//多核系能测试
package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	func3()
}

func func1() {
	t1 := time.Now()
	count := 0
	for i := 0; i < 9000000000; i++ {
		count = count + i
	}
	t2 := time.Now()
	fmt.Printf("cost:%d,count:%d\n", t2.Sub(t1), count)
}

func func2() {
	t1 := time.Now()
	count := 0
	for i := 0; i < 9000000000; i++ {
		count += i

	}
	t2 := time.Now()
	fmt.Printf("cost:%d,count:%d\n", t2.Sub(t1), count)
}

func func3() {
	t1 := time.Now()
	NCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(NCPU)
	chs := make([]chan int, NCPU)
	for i := 0; i < NCPU; i++ {
		chs[i] = make(chan int)
		n := 9000000000 / NCPU
		go run(i*n, (i+1)*n, chs[i])
	}

	count := 0
	for i := 0; i < NCPU; i++ {
		t := <-chs[i]
		count = count + t
	}
	t2 := time.Now()

	fmt.Printf("cpu num:%d,cost:%d,count:%d\n", NCPU, t2.Sub(t1), count)
}

func run(i, n int, ch chan int) {
	count := 0
	for i := i; i < n; i++ {
		count = count + i
	}
	ch <- count

}
