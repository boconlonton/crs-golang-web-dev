package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", setCookie)
	http.HandleFunc("/read", readCookie)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func setCookie(res http.ResponseWriter, req *http.Request) {
	http.SetCookie(res, &http.Cookie{
		Name:  "my-cookie",
		Value: "some value",
		Path:  "/",
	})
	fmt.Fprintln(res, "COOKIE WRITTEN - CHECK YOUR BROWSER")
	fmt.Fprintln(res, "in chrome go to: dev tools / application / cookies")
}

func readCookie(res http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("my-cookie")
	if err != nil {
		http.Error(res, http.StatusText(400), http.StatusBadRequest)
		return
	}

	fmt.Fprintln(res, "YOUR COOKIE:", c)
}
