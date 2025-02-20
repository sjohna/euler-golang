package euler5

import (
	"euler/utilities"
)

/*
https://projecteuler.net/problem=5

What is the smallest positive number that is evenly divisible by all the numbers from 1 to 20?
*/

func LoopLCMFunction() int {
	lcm := 1

	for i := 2; i <= 20; i++ {
		lcm = utilities.LCM(lcm, i)
	}

	return lcm
}

func GeneratorReduce() int {
	return utilities.NaturalNumbers().TakeWhile(func(n int) bool { return n <= 20 }).Reduce(utilities.LCM, 1)
}
