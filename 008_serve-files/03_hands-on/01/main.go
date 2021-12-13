package main

import (
	"io"
	"log"
	"net/http"
	"text/template"
)

var tpl *template.Template

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/toby.jpg", dogPic)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func foo(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "foo ran")
}

func dog(res http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("dog.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	data := struct {
		FileName string
	}{
		FileName: "/toby.jpg",
	}
	tpl.ExecuteTemplate(res, "dog.gohtml", data)
}

func dogPic(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "toby.jpg")
}
