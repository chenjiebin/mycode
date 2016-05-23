package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now().Unix())
}

func getNow() {

	// 获取当前时间
	fmt.Println(time.Now())
	// output: 2016-03-09 09:47:49.900023477 +0800 CST

	// 获取当前时间戳
	fmt.Println(time.Now().Unix())
	// output: 1457488069
}

func getHourTimestamp(timestamp int64) {
	hour := time.Unix(timestamp, 0).Format("2006-01-02 15") + ":00:00"
	fmt.Println(hour)
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", hour, time.Local)
	fmt.Println(t.Unix())
}

// 格式化时间显示
func formatTime() {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))

	t := time.Now().Unix()
	fmt.Println(t)
	fmt.Println(time.Unix(t, 0).String())
}

// 获取年份
func getYear() {
	//t := time.Now()
	//fmt.Println(t.Year())

	t := time.Unix(469296000, 0)
	fmt.Println(t.Year())
}

func getMonth() {
	m := time.Now().Month()
	fmt.Println(m, int(m))
}

//获取今天开始时间戳
func getTodayStartTimeUnix() int64 {
	t, _ := time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02")+" 00:00:00")
	fmt.Println(t.Unix() - 8*3600)
	return t.Unix() - 8*3600
}

func getDayStart(ts int64) {
	t := time.Unix(ts, 0).Format("2006-01-02")
	sts, err := time.Parse("2006-01-02", t)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sts.Unix())
}
