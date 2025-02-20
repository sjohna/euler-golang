package euler3

import (
	"euler/utilities"
)

/*
https://projecteuler.net/problem=3

What is the largest prime factor of 600851475143?
*/

func GeneratorWithLoop() int {
	curr := 600851475143
	nextPrime := utilities.PrimeGenerator()
	var currPrime int

	for curr > 1 {
		currPrime = nextPrime()
		for curr%currPrime == 0 {
			curr /= currPrime
		}
	}

	return currPrime
}
