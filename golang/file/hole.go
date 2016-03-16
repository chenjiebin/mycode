// 存在文件空洞的操作
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	writeHole()
	//	readHolePerformance()
}

func writeHole() {
	file, err := os.OpenFile("hole", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	defer file.Close()

	b := []byte{1}
	n, err := file.WriteAt(b, 6000)
	fmt.Println(n, err)
	// output: 1 <nil>
}

func readHole() {
	file, _ := os.OpenFile("hole", os.O_RDONLY, 0600)
	defer file.Close()

	//	b := make([]byte, 1)
	//	n, err := file.ReadAt(b, 512)
	//	fmt.Println(b, n, err)

	//	b := make([]byte, 1)
	//	n, err := file.ReadAt(b, 1024)
	//	fmt.Println(b, n, err)

	for i := 0; i <= 1024; i++ {
		b := make([]byte, 1)
		n, err := file.ReadAt(b, int64(i))
		fmt.Println(b, n, err)
	}
}

func readHolePerformance() {
	file, _ := os.OpenFile("hole", os.O_RDONLY, 0600)
	defer file.Close()

	fi, _ := file.Stat()
	fmt.Println(fi.Size())

	t0 := time.Now()

	//	for i := int64(0); i < fi.Size(); i++ {
	//		b := make([]byte, 1)
	//		n, err := file.ReadAt(b, int64(i))
	//		fmt.Println(b, n, err)
	//	}
	b := make([]byte, fi.Size())
	file.Read(b)

	count := 0
	for _, v := range b {
		if v == 1 {
			count++
		}
	}
	fmt.Println(count)

	t1 := time.Now()
	fmt.Println(t1.Sub(t0))
}
