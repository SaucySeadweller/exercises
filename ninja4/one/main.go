package main

import "fmt"

func main() {
	x := [6]int{1, 2, 4, 8, 16, 32}
	for i, l := range x {
		fmt.Println(i, l)
	}
	fmt.Printf("%T\n", x)
}
