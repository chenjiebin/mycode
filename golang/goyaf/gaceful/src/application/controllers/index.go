package controllers

import (
	"git.oschina.net/iceup/goyaf"
)

type Index struct {
	Base
}

func (this *Index) IndexAction() {
	var body string
	body = "hello world"
	this.GetResponse().AppendBody(body)
}

func init() {
	goyaf.Debug("init controller index")
}
