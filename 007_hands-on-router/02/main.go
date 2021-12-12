package main

import (
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/me", meHandler)
	http.HandleFunc("/dog", dogHandler)
	http.ListenAndServe(":8080", nil)
}

func dogHandler(res http.ResponseWriter, req *http.Request) {
	data := struct {
		PageName string
		Context  string
	}{
		"Dog",
		"dog dog dog",
	}
	tpl.ExecuteTemplate(res, "dog.gohtml", data)
}

func indexHandler(res http.ResponseWriter, req *http.Request) {
	data := struct {
		PageName string
		Context  string
	}{
		"Home",
		"index index index",
	}
	tpl.ExecuteTemplate(res, "index.gohtml", data)
}

func meHandler(res http.ResponseWriter, req *http.Request) {
	data := struct {
		PageName string
		Context  string
	}{
		"Me",
		"me me me",
	}
	tpl.ExecuteTemplate(res, "me.gohtml", data)
}
