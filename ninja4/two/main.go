package main

import "fmt"

func main() {
	x := []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024}
	for i, l := range x {
		fmt.Println(i, l)
	}
	fmt.Printf("%T\n", x)
}
