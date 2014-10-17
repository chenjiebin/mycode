package controllers

type Index struct {
	Base
}

func (this *Index) IndexAction() {
	this.GetResponse().SetBody("index index index")
}
