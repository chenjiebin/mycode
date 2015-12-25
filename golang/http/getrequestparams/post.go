package main

import (
	"fmt"
	"log"
	//"mime"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {

		}
		fmt.Println("body", string(body))
		fmt.Fprintln(w, string(body))
		//		r.ParseMultipartForm(32 << 20)
		//		fmt.Fprintln(w, r.MultipartForm)
		//		if r.MultipartForm != nil {
		//			values := r.MultipartForm.Value["id"]
		//			if len(values) > 0 {
		//				fmt.Fprintln(w, values[0])
		//			}
		//		}

		//ct, _, _ := mime.ParseMediaType(r.Header.Get("Content-Type"))
		//switch ct {
		//case "application/x-www-form-urlencoded":
		//	fmt.Fprintln(w, r.PostFormValue("id"))
		//case "multipart/form-data":
		//	r.ParseMultipartForm(32 << 20)
		//	fmt.Fprintln(w, r.MultipartForm)
		//	values := r.MultipartForm.Value["id"]
		//	if len(values) > 0 {
		//		fmt.Fprintln(w, values[0])
		//	}
		//}
	})

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
