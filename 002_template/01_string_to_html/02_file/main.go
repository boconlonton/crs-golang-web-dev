package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := "bocon"
	str := fmt.Sprint(`
<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Hello World!</title>
</head>
<body>
<h1>` + name + `</h1>
</body>
</html>
`)
	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal("error creating file", err)
	}
	defer nf.Close()

	_, err = io.Copy(nf, strings.NewReader(str))
	if err != nil {
		log.Fatal("error creating file", err)
	}
}
