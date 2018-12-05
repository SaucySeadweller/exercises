package main

import "fmt"

func main() {
	m := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}
	fmt.Println(m)
	for r, t := range m {
		fmt.Println("This is the record for", r)
		for i, t2 := range t {
			fmt.Println("\t", i, t2)
		}
	}
}
