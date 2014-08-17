package controllers

import (
	"fmt"
	"goyaf"
	"html/template"
	"net/http"
)

type Index struct {
	goyaf.Controller
}

func (this *Index) IndexAction(w http.ResponseWriter, r *http.Request) {

	t, _ := template.ParseFiles("/Users/chenjiebin/Sites/mycode/golang/project/webframework/views/index/index.html")
	t.Execute(w, nil)
}

func (this *Index) TestAction() {
	fmt.Println("hehe")
	//fmt.Fprintln(this.Controller.ResponseWriter, "hehe")
}
