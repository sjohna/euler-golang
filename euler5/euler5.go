package main

import (
	"euler"
	"fmt"
)

func main() {
	lcm := 1

	for i := 2; i <= 20; i++ {
		lcm = euler.LCM(lcm, i)
	}

	fmt.Println(lcm)
}
