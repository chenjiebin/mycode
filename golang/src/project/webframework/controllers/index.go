package controllers

import (
	"fmt"
	"goyaf"
	//"html/template"
)

type Index struct {
	goyaf.GoyafController
	Path string
}

func (this *Index) IndexAction() {
	id := this.GetRequest().GetQuery("id", "0")
	name := this.GetRequest().GetPost("name", "iceup")

	fmt.Println(id, name)

	//t, _ := template.ParseFiles("/Users/chenjiebin/Sites/mycode/golang/project/webframework/views/index/index.html")
	//t.Execute(w, nil)
}

func (this *Index) TestAction() {
	fmt.Println("hehe")
	//fmt.Fprintln(this.Controller.ResponseWriter, "hehe")
}
