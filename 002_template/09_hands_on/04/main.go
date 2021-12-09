package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type item struct {
	Name string
}

type meal struct {
	Term       string
	Components []item
}

type restaurant struct {
	Name string
	Menu []meal
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	menuDeporte := []meal{
		{
			"Breakfast",
			[]item{
				{"Spring Rolls"},
				{"French Onion Soup"},
				{"Caesar Salad"},
			},
		},
		{
			"Lunch",
			[]item{
				{"Mixed Green Salad"},
				{"Garden Vegetables"},
				{"French Fries"},
			},
		},
		{
			"Dinner",
			[]item{
				{"Apple Pie with Cream"},
				{"Fruit Salad"},
				{"Vanilla Ice Cream"},
			},
		},
	}

	menuLeMoon := []meal{
		{
			"Breakfast",
			[]item{
				{"Spring Rolls"},
				{"French Onion Soup"},
				{"Caesar Salad"},
			},
		},
		{
			"Lunch",
			[]item{
				{"Mixed Green Salad"},
				{"Garden Vegetables"},
				{"French Fries"},
			},
		},
		{
			"Dinner",
			[]item{
				{"Apple Pie with Cream"},
				{"Fruit Salad"},
				{"Vanilla Ice Cream"},
			},
		},
	}

	menuDiePie := []meal{
		{
			"Breakfast",
			[]item{
				{"Spring Rolls"},
				{"French Onion Soup"},
				{"Caesar Salad"},
			},
		},
		{
			"Lunch",
			[]item{
				{"Mixed Green Salad"},
				{"Garden Vegetables"},
				{"French Fries"},
			},
		},
		{
			"Dinner",
			[]item{
				{"Apple Pie with Cream"},
				{"Fruit Salad"},
				{"Vanilla Ice Cream"},
			},
		},
	}

	restaurants := []restaurant{
		{"De Porte", menuDeporte},
		{"Le Moon", menuLeMoon},
		{"Die Pie", menuDiePie},
	}
	err := tpl.Execute(os.Stdout, restaurants)
	if err != nil {
		log.Fatalln(err)
	}
}
