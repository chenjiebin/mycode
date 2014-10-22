package controllers

type Index struct {
	Base
}

func (this *Index) IndexAction() {
	var body string
	body = "index index index"

	id := this.GetRequest().GetParam("id", "0")
	body = body + "<br />params id:" + id

	//params := this.GetRequest().GetParams()

	this.GetResponse().AppendBody(body)
	//this.GetResponse().AppendBody(params)
}
