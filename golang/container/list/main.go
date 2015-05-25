//偶然看到，list的一个坑m，遍历删除list元素无效
//http://www.cnblogs.com/ziyouchutuwenwu/p/3780800.html
package main

import (
	"container/list"
	"fmt"
)

func base() {
	//初始化一个list
	l := list.New()
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(4)

	fmt.Println("Before Removing...")
	//遍历list，删除元素
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println("removing", e.Value)
		l.Remove(e)
	}
	fmt.Println("After Removing...")
	//遍历删除完元素后的list
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

type User struct {
	Name string
}

func pointEntry() {
	user := &User{}
	l := list.New()
	l.PushBack(*user)

	user.Name = "陈杰斌"

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

func main() {
	pointEntry()
}

//package main

//import (
//    "container/list"
//    "fmt"
//)

//func main() {
//    l := list.New()
//    l.PushBack(1)
//    l.PushBack("asd")
//    l.PushBack(3)
//    l.PushBack(4)
//    fmt.Println("Before Removing...")
//    var n *list.Element
//    for e := l.Front(); e != nil; e = n {
//        fmt.Println("removing", e.Value)
//        n = e.Next()
//        l.Remove(e)
//    }
//    fmt.Println("After Removing...")
//    for e := l.Front(); e != nil; e = e.Next() {
//        fmt.Println(e.Value)
//    }
//}
