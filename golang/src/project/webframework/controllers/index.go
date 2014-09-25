package controllers

type Index struct {
	Abstract
}

func (this *Index) IndexAction() {
	this.GetResponse().SetBody("index index index")
}
