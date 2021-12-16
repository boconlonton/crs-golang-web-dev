package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {

	// body
	bs := make([]byte, req.ContentLength)
	req.Body.Read(bs)
	body := string(bs)

	err := tpl.ExecuteTemplate(w, "index.gohtml", body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatalln(err)
	}

	if req.Method == http.MethodPost {
		// Read
		data, _, err := req.FormFile("upload")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		bss, err := ioutil.ReadAll(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Println(string(bss))
	}
}
