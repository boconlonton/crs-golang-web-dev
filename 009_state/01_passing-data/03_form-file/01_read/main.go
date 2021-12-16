package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func foo(res http.ResponseWriter, req *http.Request) {
	var s string
	fmt.Println(req.Method)
	if req.Method == http.MethodPost {
		// Open
		f, h, err := req.FormFile("q")
		if err != nil {
			fmt.Println("here")
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()

		// Print out the information
		fmt.Println("\nFile:", "\nheader:", h, "\nerr", err)

		// Read
		bs, err := ioutil.ReadAll(f)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		s = string(bs)
	}
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(res, `
	<form method="POST" enctype="multipart/form-data">
	<input type="file" name="q">
	<input type="submit">
	</form>
	<br>`+s)
}
