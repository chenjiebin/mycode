/**
 * 从服务器获取系统时间解决时间差的问题
 */

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func getSystemTime() int64 {
	resp, err := http.Get("http://app.pba.cn/test/time.php")
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	systemTime, err := strconv.ParseInt(string(body), 10, 64)

	return systemTime
}

func main() {
	startTime := time.Now().UnixNano()
	fmt.Println(startTime)

	systemTime := getSystemTime()

	endTime := time.Now().UnixNano()

	//systemTime = systemTime + endTime - startTime

	fmt.Println(systemTime)
	fmt.Println(endTime)
}
