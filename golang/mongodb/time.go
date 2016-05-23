// 演示一下mongodb中的cst和utc时间戳
package main

import (
	"fmt"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// get mongodb db
func getDB() *mgo.Database {
	session, err := mgo.Dial("192.168.1.119:27017")
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)
	db := session.DB("meiqu618_20150211")
	return db
}

func main() {
	db := getDB()
	c := db.C("clients")

	//	{
	//	    "_id" : ObjectId("559ced6a0651f3a66f6b9426"),
	//	    "updated_at" : ISODate("2015-07-08T09:29:14.002Z")
	//	}
	type Client struct {
		Id        bson.ObjectId `bson:"_id,omitempty"` // id
		UpdatedAt time.Time     `bson:"updated_at"`    // 最后更新时间
	}
	client := Client{}
	err := c.Find(bson.M{"_id": bson.ObjectIdHex("559ced6a0651f3a66f6b9426")}).One(&client)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", client)
	// output: {Id:ObjectIdHex("559ced6a0651f3a66f6b9426") UpdatedAt:2015-07-08 17:29:14.002 +0800 CST}

	timeStr := "2015-07-08 17:29:14.002 +0800 CST"
	t, _ := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", timeStr)
	fmt.Printf("%+v\n", t)

	err = c.Find(bson.M{
		"_id":        bson.ObjectIdHex("559ced6a0651f3a66f6b9426"),
		"updated_at": bson.M{"$gte": t},
	}).One(&client)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", client)
}
