package controllers

import (
	"fmt"
	"html/template"
	"net/http"
)

type Index struct {
}

func (Index *Index) IndexAction(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("views/index/index.html")

	t.Execute(w, nil)
}

func (Index *Index) TestAction(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello test!")
}
