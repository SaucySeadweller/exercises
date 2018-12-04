package main

import (
	"math/rand"
	"time"
)

var powerlevel int

func random(min, max int) int {
	powerlevel := rand.Intn(max - min)
	return powerlevel
}
func main() {
	rand.Seed(time.Now().UnixNano())
	powerlevel := random(1, 15000)
	if powerlevel > 9000 {
		println(powerlevel, "It's over 9000!!!")
	} else if powerlevel == 42 {
		println(powerlevel, "Perfection")
	} else {
		println(powerlevel, "What a pathetic powerlevel.")
	}
}
