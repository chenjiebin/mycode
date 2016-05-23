// 测试mongodb连接问题
package main

import (
	"fmt"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	fmt.Println("start session")
	session, err := mgo.Dial("172.16.27.134:10001")
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	//	time.Sleep(10 * 1e9)
	fmt.Println("start session copy")
	s := session.Copy()
	//	time.Sleep(10 * 1e9)
	db := s.DB("test")

	time.Sleep(3 * 1e9)
	fmt.Println("start find")

	for i := 0; i < 10000; i++ {
		go func() {
			c := db.C("user")
			type User struct {
				Name string "bson:`name`"
				Age  int    "bson:`age`"
			}
			users := []User{}
			err = c.Find(bson.M{}).All(&users)
			if err != nil {
				panic(err)
			}
			fmt.Println(users)
		}()
	}
	time.Sleep(10 * 1e9)

	fmt.Println("s close")
	s.Close()
	fmt.Println("session close")
	session.Close()

	time.Sleep(10 * 1e9)
}
