package controllers

import (
	"git.oschina.net/iceup/goyaf"
)

type Index struct {
	Base
}

func (this *Index) IndexAction() {
	var body string
	body = "index index index"

	id := this.GetRequest().GetParam("id", "0")
	body = body + "<br />params id:" + id

	this.GetResponse().AppendBody(body)
}

func init() {
	goyaf.Debug("init controller index")
}
