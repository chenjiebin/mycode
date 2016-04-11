package main

import (
	"os"
	"text/template"
)

func main() {
	type Page struct {
		Title string
	}
	page := Page{"test page"}
	tmpl, err := template.New("test").Parse(`<!DOCTYPE HTML>
<html>
<head>
<meta charset="UTF-8" />
<title>{{.Title}}</title>
</head>
<body>



</body>
</html>`)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, page)
	if err != nil {
		panic(err)
	}
}
