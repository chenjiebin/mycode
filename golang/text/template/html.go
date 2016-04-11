//
package main

import (
	"html/template"
	"net/http"
)

func main() {
	t, _ := template.New("fieldname example").Parse("hello {{.UserName}}!")
	p := Person{UserName: "Astaxie"}
	t.Execute(os.Stdout, p)
}
