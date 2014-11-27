package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

func main() {
	fmt.Println("aaa")

	conn, err := redis.DialTimeout("tcp", "192.168.3.233:6379", 0, 1*time.Second, 1*time.Second)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(conn.Do("select", "15"))
	fmt.Println(conn.Do("set", "name", "cjb"))

	fmt.Println(conn)
	time.Sleep(30 * 1e9)
	conn.Close()
}
