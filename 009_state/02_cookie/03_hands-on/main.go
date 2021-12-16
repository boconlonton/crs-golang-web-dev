package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("counter")

	if err != nil {
		if err == http.ErrNoCookie {
			c = &http.Cookie{
				Name:  "counter",
				Value: "0",
				Path:  "/",
			}
		} else {
			log.Fatalln(err)
		}
	}
	count, err := strconv.Atoi(c.Value)
	count++
	c.Value = strconv.Itoa(count)
	http.SetCookie(res, c)
	io.WriteString(res, c.Value)
}
