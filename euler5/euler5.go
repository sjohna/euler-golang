package euler5

import (
	"euler/utilities"
	"fmt"
)

/*
https://projecteuler.net/problem=5

What is the smallest positive number that is evenly divisible by all the numbers from 1 to 20?
*/

func Euler5() {
	lcm := 1

	for i := 2; i <= 20; i++ {
		lcm = utilities.LCM(lcm, i)
	}

	fmt.Println(lcm)
}
