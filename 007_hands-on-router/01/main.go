package main

import (
	"io"
	"net/http"
)

func dogHandler(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog dog")
}

func indexHandler(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "index index index")
}

func meHandler(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "It's me! Toommyyyy!")
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/me", meHandler)
	http.HandleFunc("/dog", dogHandler)
	http.ListenAndServe(":8080", nil)
}
