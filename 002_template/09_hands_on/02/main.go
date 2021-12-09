package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	data := []struct {
		Name    string
		Address string
		City    string
		Zip     string
		Region  string
	}{
		{
			Name:    "Hotel 1",
			Address: "01 Avenue",
			City:    "California",
			Zip:     "70001",
			Region:  "Southern",
		},
		{
			Name:    "Hotel 2",
			Address: "02 Avenue",
			City:    "California",
			Zip:     "70001",
			Region:  "Central",
		},
		{
			Name:    "Hotel 3",
			Address: "01 Avenue",
			City:    "California",
			Zip:     "70001",
			Region:  "Northern",
		},
		{
			Name:    "Hotel 4",
			Address: "01 Avenue",
			City:    "California",
			Zip:     "70001",
			Region:  "Southern",
		},
		{
			Name:    "Hotel 5",
			Address: "01 Avenue",
			City:    "California",
			Zip:     "70001",
			Region:  "Central",
		},
	}
	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
