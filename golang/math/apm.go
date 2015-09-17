//十亿个数字中挑选出最小的10个
//思路来源：http://www.oschina.net/code/snippet_2349747_50840

package main

import (
	"fmt"
	//	"math"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU() - 4)
	//开始时间
	st := time.Now().UnixNano()

	result := make([]int, 10)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100000000; i++ {
		randNum := rand.Intn(100000000)
		if i < 10 {
			result[i] = randNum
			continue
		}

		maxNumIndex := 0
		for k, v := range result {
			if k == 0 {
				continue
			}
			if v > result[maxNumIndex] {
				maxNumIndex = k
			}
		}

		if randNum < result[maxNumIndex] {
			result[maxNumIndex] = randNum
		}
	}

	//结束时间
	et := time.Now().UnixNano()
	fmt.Println(et-st, "纳秒")
	fmt.Println(result)

}
