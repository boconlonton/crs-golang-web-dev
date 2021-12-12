package main

import (
	"fmt"
	"net/http"
	"text/template"
)
import "github.com/julienschmidt/httprouter"

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	mux := httprouter.New()
	mux.GET("/user", user)

	http.ListenAndServe(":8080", mux)
}

func user(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "USER, %s!\n", ps.ByName("user"))
}
