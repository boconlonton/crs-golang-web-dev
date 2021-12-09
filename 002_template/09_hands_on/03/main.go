package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type dish struct {
	Name string
}

type menu struct {
	Term       string
	Components []dish
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	data := []menu{
		{
			"Breakfast",
			[]dish{
				{"Spring Rolls"},
				{"French Onion Soup"},
				{"Caesar Salad"},
			},
		},
		{
			"Lunch",
			[]dish{
				{"Mixed Green Salad"},
				{"Garden Vegetables"},
				{"French Fries"},
			},
		},
		{
			"Dinner",
			[]dish{
				{"Apple Pie with Cream"},
				{"Fruit Salad"},
				{"Vanilla Ice Cream"},
			},
		},
	}
	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
