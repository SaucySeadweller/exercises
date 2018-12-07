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
			"Rocky Road,",
			"Chocolate,",
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

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for _, v := range m {
		fmt.Println(v.first)
		fmt.Println(v.last)
		for i, val := range v.icecreamFlavors {
			fmt.Println(i, val)
		}
		fmt.Println("")
	}
}
