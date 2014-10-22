package main

import (
	"fmt"
	"time"
)

func main() {
	getTodayStartTimeUnix()
}

func formatTime() {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))

	t := time.Now().Unix()
	fmt.Println(t)
	fmt.Println(time.Unix(t, 0).String())
}

func getYear() {
	//t := time.Now()
	//fmt.Println(t.Year())

	t := time.Unix(469296000, 0)
	fmt.Println(t.Year())
}

//获取今天开始时间戳
func getTodayStartTimeUnix() int64 {
	t, _ := time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02")+" 00:00:00")
	fmt.Println(t.Unix() - 8*3600)
	return t.Unix() - 8*3600
}
