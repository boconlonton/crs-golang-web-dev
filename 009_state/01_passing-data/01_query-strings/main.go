package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func foo(res http.ResponseWriter, req *http.Request) {
	v := req.FormValue("q")
	fmt.Fprintln(res, "Do my search: "+v)
}
