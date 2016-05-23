package main

import (
	"fmt"
	"log"
	"net/http"

	"gopkg.in/mgo.v2"
)

var session *mgo.Session

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	db := session.DB("test")

	for i := 0; i < 100; i++ {
		c := db.C("user")
		type User struct {
			Name string "bson:`name`"
			Age  int    "bson:`age`"
		}
		users := []User{}
		err := c.Find(nil).All(&users)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(users)
	}

	fmt.Fprintf(w, "1")

}

func main() {
	http.HandleFunc("/", sayhelloName)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func init() {
	session, _ = mgo.Dial("172.16.27.134:10001")
	session.SetMode(mgo.Monotonic, true)
	fmt.Println("init mongodb	")
}
