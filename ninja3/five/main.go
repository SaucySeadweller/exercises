package main

import "fmt"

func main() {

	for i := 10; i <= 100; i++ {
		fmt.Printf("%v  %v", i, i%4)
	}
}
