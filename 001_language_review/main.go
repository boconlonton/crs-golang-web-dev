package main

import "fmt"

type person struct {
	fname string
	lname string
}

type secretAgent struct {
	person
	licenseToKill bool
}

func (p person) speak() {
	fmt.Println(p.fname, `says "Good morning, James."`)
}

func (sa secretAgent) speak() {
	fmt.Println(sa.fname, sa.lname, `says, "Shaken, not stirred"`)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	// Slice
	xi := []int{2, 4, 7, 9, 42}
	fmt.Println(xi)

	// Map
	m := map[string]int{
		"Todd": 25,
		"Joe":  42,
	}
	fmt.Println(m)

	// Struct
	p1 := person{
		"Miss",
		"Moneypenny",
	}
	fmt.Println(p1)

	// Methods
	p1.speak()

	// Embedded struct (Inheritance)
	sa1 := secretAgent{
		person{
			"James",
			"Bond",
		},
		true,
	}
	sa1.speak()

	// Interface (Polymorphism)
	saySomething(p1)
	saySomething(sa1)

}
