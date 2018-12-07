package main

import "fmt"

func main() {
	cat := struct {
		name  string
		color string
		breed string
	}{

		name:  "Smitty",
		color: "Calico",
		breed: "Norwegian Ridgeback",
	}
	fmt.Println(cat)
}
