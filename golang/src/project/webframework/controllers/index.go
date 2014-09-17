package controllers

import (
	"fmt"
	"goyaf"
	"html/template"
	"net/http"
)

type Index struct {
	goyaf.GoyafController
	Path string
}

func (this *Index) IndexAction(w http.ResponseWriter, r *http.Request) {
	fmt.Println("index")

	//t, _ := template.ParseFiles("/Users/chenjiebin/Sites/mycode/golang/project/webframework/views/index/index.html")
	//t.Execute(w, nil)
}

func (this *Index) TestAction() {
	fmt.Println("hehe")
	//fmt.Fprintln(this.Controller.ResponseWriter, "hehe")
}
