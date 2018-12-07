package main

import "fmt"

type person struct {
	first           string
	last            string
	icecreamFlavors []string
}

func main() {
	p1 := person{
		first: "Darius",
		last:  "The Hand of Noxus",
		icecreamFlavors: []string{
			"Chocolate,",
			"Rocky Road,",
			"Tears"},
	}
	p2 := person{
		first: "Draven",
		last:  "The Glorious Executioner",
		icecreamFlavors: []string{
			"Draven,",
			"Red Velvet,",
			"Draaaaven"},
	}
	fmt.Println(p1.first)
	fmt.Println(p1.last)
	for i, v := range p1.icecreamFlavors {
		fmt.Println(i, v)
	}

	fmt.Println(p2.first)
	fmt.Println(p2.last)
	for i, v := range p2.icecreamFlavors {
		fmt.Println(i, v)
	}
}
